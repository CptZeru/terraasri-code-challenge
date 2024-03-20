package repository

import (
	"CptZeru/terraasri-code-challenge/constants"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var EmojiDictionary map[string]string
var ReversedEmojiDictionary map[string]string

// GetEmojiDictionary initiate EmojiDictionary with data from json file and return the data.
func GetEmojiDictionary() map[string]string {
	initIfEmojiDictionaryDoesntExist()
	return EmojiDictionary
}

// GetReversedEmojiDictionary initiate ReversedEmojiDictionary with data from json file and return the data.
func GetReversedEmojiDictionary() map[string]string {
	initIfReverseEmojiDictionaryDoesntExist()
	return ReversedEmojiDictionary
}

// UpdateEmojiDictionaries update emoji dictionaries in current session and in json file.
func UpdateEmojiDictionaries(emoji string, meaning string) error {
	initIfEmojiDictionaryDoesntExist()
	initIfReverseEmojiDictionaryDoesntExist()
	EmojiDictionary[emoji] = meaning
	ReversedEmojiDictionary[meaning] = emoji
	writeFile(EmojiDictionary, constants.EmojiDictionaryFN)
	writeFile(ReversedEmojiDictionary, constants.ReversedEmojiDictionaryFN)
	return nil
}

// initIfEmojiDictionaryDoesntExist initiate EmojiDictionary data from json file
func initIfEmojiDictionaryDoesntExist() {
	if len(EmojiDictionary) == 0 {
		data := openAndReadFile(constants.EmojiDictionaryFN)
		EmojiDictionary = data
	}
}

// initIfEmojiDictionaryDoesntExist initiate ReversedEmojiDictionary data from json file
func initIfReverseEmojiDictionaryDoesntExist() {
	if len(ReversedEmojiDictionary) == 0 {
		data := openAndReadFile(constants.ReversedEmojiDictionaryFN)
		ReversedEmojiDictionary = data
	}
}

// openAndReadFile open and read given fileName file while returning unmarshalled data.
func openAndReadFile(fileName string) map[string]string {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	fmt.Printf("Successfully opened %v file\n", fileName)

	byteValue, _ := io.ReadAll(jsonFile)
	var dictionary map[string]string
	json.Unmarshal([]byte(byteValue), &dictionary)

	return dictionary
}

// writeFile write json file to create / update its value.
func writeFile(dictionary map[string]string, fileName string) error {
	file, err := json.MarshalIndent(dictionary, "", "")
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, file, 0644)
	if err != nil {
		return err
	}
	return nil
}
