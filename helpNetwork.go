package goHelp

import (
	strings
	strconv
)


func checkValidIP(ip string) bool {
        if ip == "" {
                return false
        }
        checkIP := strings.Split(ip, ".")
        if len(checkIP) != 4 {
                return false
        }
        for _, ip := range checkIP {
                if octet, err := strconv.Atoi(ip); err != nil {
                        return false
                } else if octet < 0 || octet > 255 {
                        return false
                }
        }
        return true
}


