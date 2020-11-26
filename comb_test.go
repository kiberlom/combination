package combination

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	Set(10, "該命令與指定的遠程項目連接，並獲取您尚未擁有的所有項目數據。運行該命令後，您應該具有指向該遠程項目中所有分支的鏈接，可以隨時檢查或合併這些鏈接。克隆存儲庫時，clone命令會自動將該遠程存儲庫添加為“ origin”名稱。通過這種方式，git fetch origin可以在克隆服務器（或獲取更改）後獲取推送到該服務器的所有工作。重要的是要注意git fetch命令將數據提取到本地存儲庫中，但不會將其合併到您的任何工作中或修改您當前正在處理的內容。準備好後，您需要手動將這些數據與您的數據合併。")

	a := GetFirst()

	for i := 0; i < 1000000; i++ {
		s, err := Next(a)

		if err != nil {
			fmt.Println(err)
		}

		a = s

		fmt.Println(s)
	}
}
