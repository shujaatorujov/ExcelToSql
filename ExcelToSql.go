package main

import (
	"bufio"
	"flag"
	"github.com/google/uuid"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"strings"
)

func main() {
	var sheets [][][]string
	var displayTextScript []string
	var dictItemScript []string
	var err error

	// Flags
	fileName := flag.String("f", "dictionary.xlsx", "xlsx file path")
	dictName := flag.String("dn", "dictionaryName", "dictionary name")
	dictCode := flag.String("dc", "dictionaryCode", "dictionary code")
	flag.Parse()

	sheets, err = xlsx.FileToSlice(*fileName)
	if err != nil {
		log.Fatal("File not found")
	}
	displayTextScript = []string{}
	dictItemScript = []string{}
	for r := 1; r < len(sheets[0]); r++ {
		dtId := uuid.New().String()
		diId := uuid.New().String()
		displayText := "insert into common.display_text (id, column_name, ru, en, az) values (" + "'" + dtId + "'," + "'dictionary.dictionary_item.title_ref_id'"
		dictItem := "insert into dictionary.dictionary_item (id, dictionary_name, code, title_ref_id, avis_code, parent_dictionary_item_id) VALUES (" + "'" + diId + "'," + "'" + *dictName + "'," + "'" + *dictCode + "'," + "'" + dtId + "'," + "'" + sheets[0][r][0] + "'," + "null);"
		for c := 1; c < len(sheets[0][r]); c++ {
			evaluatedStr := strings.Trim(strings.ReplaceAll(sheets[0][r][c], "'", "''"), " \n")
			displayText = displayText + ",'" + evaluatedStr + "'"
		}
		displayText = displayText + ");"
		displayTextScript = append(displayTextScript, displayText)
		dictItemScript = append(dictItemScript, dictItem)
	}
	//fmt.Println(displayTextScript)
	createOrUpdateFile(displayTextScript, "displayText.sql")
	createOrUpdateFile(dictItemScript, "dictItem.sql")
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
