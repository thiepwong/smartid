// luhn luhn
// generate the code base on Luhn algorithm
// Author: Thiep Wong
// Created: 13.3.2019
package luhn

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/thiepwong/smartid/pkg/logger"
)

// Valid returns a boolean indicating if the argument was valid according to the Luhn algorithm.
func Valid(luhnString string) bool {
	checksumMod := calculateChecksum(luhnString, false) % 10

	return checksumMod == 0
}

// Generate creates and returns a string of the length of the argument targetSize.
// The returned string is valid according to the Luhn algorithm.
func Generate(size int) string {
	random := randomString(size - 1)
	controlDigit := strconv.Itoa(generateControlDigit(random))

	return random + controlDigit
}

// GenerateWithPrefix creates and returns a string of the length of the argument targetSize
// but prefixed with the second argument.
// The returned string is valid according to the Luhn algorithm.
func GenerateWithPrefix(size int, prefix string) string {
	size = size - 1 - len(prefix)

	random := prefix + randomString(size)
	controlDigit := strconv.Itoa(generateControlDigit(random))

	return random + controlDigit
}

func randomString(size int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	source := make([]int, size)

	for i := 0; i < size; i++ {
		source[i] = rand.Intn(9)
	}

	return integersToString(source)
}

func generateControlDigit(luhnString string) int {
	controlDigit := calculateChecksum(luhnString, true) % 10

	if controlDigit != 0 {
		controlDigit = 10 - controlDigit
	}

	return controlDigit
}

func calculateChecksum(luhnString string, double bool) int {
	source := strings.Split(luhnString, "")
	checksum := 0

	for i := len(source) - 1; i > -1; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)

		if double {
			n = n * 2
		}
		double = !double

		if n >= 10 {
			n = n - 9
		}

		checksum += n
	}

	return checksum
}

func integersToString(integers []int) string {
	result := make([]string, len(integers))

	for i, number := range integers {
		result[i] = strconv.Itoa(number)
	}

	return strings.Join(result, "")
}

// GenerateSmartID function Generate an account id base in Luhn algorithm
func GenerateSmartID(systemCode int, nodeCode int, size int) (id uint64, err error) {
	if (systemCode > 9 || systemCode < 1) || (nodeCode < 0 || nodeCode > 255) {
		fmt.Println("Loi roi")
		err = errors.New("System Code is invalid")
		return 0, err
	}

	var nodeCodeStr string
	if nodeCode >= 0 && nodeCode <= 9 {
		nodeCodeStr = "00" + strconv.Itoa(nodeCode)
	}
	if nodeCode >= 10 && nodeCode <= 99 {
		nodeCodeStr = "0" + strconv.Itoa(nodeCode)
	}

	if nodeCode > 99 {
		nodeCodeStr = strconv.Itoa(nodeCode)
	}
	randomCode := randomString(size - 6)
	accountSeed := strconv.Itoa(systemCode) + nodeCodeStr + strconv.Itoa(generateControlDigit(randomCode)) + randomCode
	accountSeed += strconv.Itoa(generateControlDigit(accountSeed))
	fmt.Println(accountSeed)
	accountID, err := strconv.ParseUint(accountSeed, 10, 64)
	if err != nil {
		logger.LogErr.Printf("Error when pass the smart id %s", err)
		return 0, err
	}
	return accountID, err
}

// func main() {

// 	x := randomString(5)
// 	fmt.Println(x)
// 	controlDigit := strconv.Itoa(generateControlDigit(x))
// 	x += controlDigit
// 	fmt.Println("So check", controlDigit)
// 	fmt.Println(x)
// 	fmt.Println(Valid(x))
// 	id, _err := GenerateSmartID(6, 0xff, 0x10)
// 	if _err != nil {
// 		logger.LogErr.Println("Da bi loi khi tao ID")
// 	}
// 	fmt.Println("Da tao tai khoan la: ", id)
// 	logger.LogInfo.Println(fmt.Sprintf("So ngau nhien: %s -------- So kiem tra: %s ------------- Ma ID da sinh: %d", x, controlDigit, id))
// }
