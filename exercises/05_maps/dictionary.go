package main

import (
    "fmt"
)

type Dictionary map[string]string

const ErrorNotFound = DictionaryError("WORD NOT FOUND")
const ErrorWordExist = DictionaryError("WORD EXIST")

type DictionaryError string

func (error DictionaryError) Error() string {
    return string(error)
}

func (dictionary Dictionary) Search(key string) (string, error) {
    var result, status = dictionary[key]
    if !status {
        return "", ErrorNotFound
    }
    return result, nil
}

func (dictionary Dictionary) Add(key string, value string) error {
    _, error := dictionary.Search(key)

    switch error {
    case ErrorNotFound:
            dictionary[key] = value
    case nil:
        return ErrorWordExist
    default:
        return error
    }
    return ErrorNotFound
}

func (dictionary Dictionary) Update(key string, value string) error {
    _, error := dictionary.Search(key)

    switch error {
    case ErrorNotFound:
        return ErrorNotFound
    case nil:
        dictionary[key] = value
    default:
        return error
    }

    return nil
}

func (dictionary Dictionary) Delete(key string) error {
    _, error := dictionary.Search(key)

    switch error {
    case ErrorNotFound:
        return ErrorNotFound;
    case nil:
        delete(dictionary, key)
    default:
        return error
    }

    return nil
}

func main() {
    var test = make(map[string]string)
    test["keyTest"] = "valueTest"
    test["name"] = "testName"
    fmt.Println(test)
}
