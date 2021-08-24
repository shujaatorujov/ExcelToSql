package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"strings"
)

// StringSlice is string slice
type StringSlice []string

func (ml *StringSlice) String() string {
	return fmt.Sprintln(*ml)
}

// Set string value in StringSlice
func (ml *StringSlice) Set(s string) error {
	*ml = append(*ml, s)
	return nil
}

func main() {
	var sheets [][][]string
	var displayTextScript []string
	var dictItemScript []string
	var err error

	// Flags
	var files StringSlice
	var dictNames StringSlice
	var dictCodes StringSlice
	flag.Var(&files, "f", "file")
	flag.Var(&dictNames, "dn", "dictionary name")
	flag.Var(&dictCodes, "dc", "dictionary code")
	flag.Parse()

	startedMap := make(map[string]string)
	for a := 0; a < len(files); a++ {
		parentMap := startedMap
		startedMap = make(map[string]string)
		fileName := files[a]
		dictCode := dictCodes[a]
		dictName := dictNames[a]
		sheets, err = xlsx.FileToSlice(fileName)
		if err != nil {
			log.Fatal("File not found")
		}
		displayTextScript = []string{}
		dictItemScript = []string{}
		for r := 1; r < len(sheets[0]); r++ {
			if sheets[0][r][0] == "" {
				break
			}
			dtId := uuid.New().String()
			diId := uuid.New().String()
			avisCode := sheets[0][r][0]
			displayText := "insert into common.display_text (id, column_name, ru, en, az) values (" + "'" + dtId + "'," + "'dictionary.dictionary_item.title_ref_id'"
			dictItem := "insert into dictionary.dictionary_item (id, dictionary_name, code, title_ref_id, avis_code, parent_dictionary_item_id) VALUES (" + "'" + diId + "'," + "'" + dictName + "'," + "'" + dictCode + "'," + "'" + dtId + "'," + "'" + avisCode + "',"
			startedMap[avisCode] = diId
			diStr := ""
			for c := 1; c < len(sheets[0][r]); c++ {
				if c < 4 {
					dtStr := strings.Trim(strings.ReplaceAll(sheets[0][r][c], "'", "''"), " \n")
					displayText = displayText + ",'" + dtStr + "'"
				} else if c == 4 {
					diStr = strings.Trim(strings.ReplaceAll(sheets[0][r][c], "'", "''"), " \n")
				}
			}
			if len(diStr) > 0 {
				if parentMap[diStr] == "" {
					dictItem = dictItem + "null"
				} else {
					dictItem = dictItem + "'" + parentMap[diStr] + "'"
				}
			} else {
				dictItem = dictItem + "null"
			}
			displayText = displayText + ");"
			dictItem = dictItem + ");"
			displayTextScript = append(displayTextScript, displayText)
			dictItemScript = append(dictItemScript, dictItem)
		}
		createOrUpdateFile(displayTextScript, dictCode+"DisplayText.sql")
		createOrUpdateFile(dictItemScript, dictCode+"DictionaryItem.sql")
	}
	fmt.Println("Dictionary files created successfully")
}

func createOrUpdateFile(texts []string, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufferedWriter := bufio.NewWriter(file)
	for _, text := range texts {
		_, err = bufferedWriter.WriteString(text + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	_ = bufferedWriter.Flush()
}
