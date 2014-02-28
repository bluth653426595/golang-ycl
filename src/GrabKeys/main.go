/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

import (
        "dlib/dbus"
        "fmt"
        "github.com/BurntSushi/xgbutil"
        "github.com/BurntSushi/xgbutil/keybind"
        "github.com/BurntSushi/xgbutil/xevent"
)

var (
        X *xgbutil.XUtil
)

func (op *Manager) RegisterAccelKey(accel string) {
        defer func() {
                if err := recover(); err != nil {
                        fmt.Println("Recover Error:", err)
                }
        }()

        grapAccelKey(X.RootWin(), accel)
}

func (op *Manager) UnregisterAccelKey(accel string) {
        defer func() {
                if err := recover(); err != nil {
                        fmt.Println("Recover Error:", err)
                }
        }()

        ungrabAccelKey(X.RootWin(), accel)
}

func main() {
        defer func() {
                if err := recover(); err != nil {
                        fmt.Println("Recover Error:", err)
                }
        }()

        var err error

        X, err = xgbutil.NewConn()
        if err != nil {
                fmt.Println("New XUtil Connection Failed:", err)
                panic(err)
        }
        keybind.Initialize(X)

        m := &Manager{}
        m.listenKeyChanged()
        dbus.InstallOnSession(m)
        dbus.DealWithUnhandledMessage()

        //m.RegisterAccelKey("mod4-p")

        xevent.Main(X)
}
