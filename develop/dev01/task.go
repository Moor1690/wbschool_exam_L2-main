package dev01

import (
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/
var mtime string = "0.beevik-ntp.pool.ntp.org" // Значение по умолчанию

// SetNTPServer позволяет установить сервер NTP
func SetNTPServer(server string) {
	mtime = server
}

// GetTime возвращает текущее время с сервера NTP
func GetTime() (time.Time, error) {
	return ntp.Time(mtime)
}
