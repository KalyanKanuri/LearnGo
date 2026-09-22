package timepkg

import (
	"fmt"
	"time"
)

func TZExecutor() {
	now := time.Now()

	barcelona, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		fmt.Println("Error loading madrid location: ", err)
		return
	}

	barcelonaTime := now.In(barcelona)

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error loading tokyo lcoation: ", err)
		return
	}

	tokyoTime := now.In(tokyo)

	fmt.Println("Current time in barcelona: ", barcelonaTime.Format(time.RFC3339))
	fmt.Println("Current time in Tokyo: ", tokyoTime.Format(time.RFC3339))
	fmt.Println("current UTC time: ", now.UTC().Format(time.RFC3339))
}
