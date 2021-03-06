/*
 * File: /store/tape_test.go                                                   *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 12:32:54                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 14:00:03                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package store

import (
	"go_start/tdd/test"
	"io/ioutil"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	file, clean := test.CreateTempFile(t, "12345")
	defer clean()

	tape := tape{file: file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)

	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
