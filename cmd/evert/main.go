package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"github.com/Ada-lave/evert-core/pkg/evert"
	"github.com/Ada-lave/evert-core/pkg/formatter"
)


func main() {
	// testZip()
	formatDoc()
	// saveNewZip("origin.docx", "origin2.docx")
}

func formatDoc() {
	file, err := os.Open("test.docx")

	if err != nil {
		panic(err)
	}

	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	evertDoc, err := evert.New(file, fileinfo.Size())
	evertFormatter := formatter.NewFormatter(evertDoc)
	evertFormatter.FormatImageBody(false)
	evertDoc.SaveFormattedDoc()
}

func testZip() {
	file, err := zip.OpenReader("generated.docx")

	if err != nil {
		panic(err)
	}

	for _, el := range file.File {
		if el.Name == "word/document.xml" {
			fmt.Println(el)
		}
	}
}

func saveNewZip(oldZipFile string, newZipFile string) {
    fileToRemove := "word/settings.xml" // Имя файла, который нужно удалить

    // Открываем старый zip-файл
    oldZip, err := zip.OpenReader(oldZipFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer oldZip.Close()

    // Создаем новый zip-файл
    newZip, err := os.Create(newZipFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer newZip.Close()

    zipWriter := zip.NewWriter(newZip)
    defer zipWriter.Close()

    // Переносим файлы из старого архива
    for _, file := range oldZip.File {
        if file.Name == fileToRemove {
            continue // Пропускаем файл, который нужно удалить
        }

        // Открываем файл из старого архива
        oldFile, err := file.Open()
        if err != nil {
            fmt.Println(err)
            continue
        }
        defer oldFile.Close()

        // Создаем запись в новом архиве
        newFileWriter, err := zipWriter.Create(file.Name)
        if err != nil {
            fmt.Println(err)
            continue
        }

        // Копируем содержимое старого файла в новый архив
        _, err = io.Copy(newFileWriter, oldFile)
        if err != nil {
            fmt.Println(err)
            continue
        }
    }

    fmt.Println("Файл", fileToRemove, "был удален из архива", oldZipFile)
}