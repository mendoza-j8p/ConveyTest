package main

import (
    "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestDivide(t *testing.T) {
    convey.Convey("Given two numbers, it should return the quotient and remainder", t, func() {
        cociente, resto := divide(17, 5)
        convey.So(cociente, convey.ShouldEqual, 3)
        convey.So(resto, convey.ShouldEqual, 2)
    })
}

func TestSum(t *testing.T) {
    convey.Convey("Given a list of numbers, it should return their sum", t, func() {
        resultado := sum(1, 2, 3, 4, 5)
        convey.So(resultado, convey.ShouldEqual, 15)
    })
}