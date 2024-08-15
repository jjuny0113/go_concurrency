package oop

import "fmt"

//type Mail struct {
//	alarm Alarm
//}
//
//type Alarm struct {
//
//}
//
//func (m *Mail) OnRecv(){
//	m.alarm.Alarm()
//}

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listeners EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listeners = listener
}

func (m *Mail) OnRecv() {
	m.listeners.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	fmt.Println("알람! 알람!")
}

var mail = &Mail{}
var listener EventListener = &Alarm{}

func Init() {
	mail.Register(listener)
	mail.OnRecv()
}
