package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var EmojiDictionary map[string]string

// initIfEmojiDictionaryDoesntExist initiate EmojiDictionary data from json file
func initIfEmojiDictionaryDoesntExist() {
	if len(EmojiDictionary) == 0 {
		data := readFile()
		EmojiDictionary = data
	}
}

// openAndReadFile open and read emoji_dictionary.json file while returning unmarshalled data.
func readFile() map[string]string {
	jsonFile, err := os.Open("./emoji_dictionary.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	fmt.Println("Successfully opened emoji_dictionary.json")

	byteValue, _ := io.ReadAll(jsonFile)
	var emojiDictionary map[string]string
	json.Unmarshal([]byte(byteValue), &emojiDictionary)

	return emojiDictionary
}

// GetEmojiDictionary initiate EmojiDictionary with data from json file and return the data.
func GetEmojiDictionary() map[string]string {
	initIfEmojiDictionaryDoesntExist()
	return EmojiDictionary
}

// UpdateEmojiDictionary update emoji dictionary in current session and in json file.
func UpdateEmojiDictionary(emoji string, meaning string) error {
	initIfEmojiDictionaryDoesntExist()
	EmojiDictionary[emoji] = meaning
	file, err := json.MarshalIndent(EmojiDictionary, "", "")
	if err != nil {
		return err
	}
	err = os.WriteFile("./emoji_dictionary.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
