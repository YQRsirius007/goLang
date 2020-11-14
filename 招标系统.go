package pos

import "fmt"

type TenderSystem struct {
	projectById    map[int][]int
	projectByPrice map[int][]int
}

func Constructor() TenderSystem {
	return TenderSystem{
		projectById:    map[int][]int{},
		projectByPrice: map[int][]int{},
	}
}

func (tenderSystem *TenderSystem) addTender(userId int, projectId int, price int) {
	userList := tenderSystem.projectById[projectId]
	if len(userList) == 0 {
		tenderSystem.projectById[projectId] = []int{}
		// tenderSystem.projectById[projectId] = append(tenderSystem.projectById[projectId], userId)
		tenderSystem.projectByPrice[projectId] = []int{}
		// tenderSystem.projectByPrice[projectId] = append(tenderSystem.projectById[projectId], price)
	}
	for i := range userList {
		if userList[i] == userId {
			return
		}
	}

	tenderSystem.projectById[projectId] = append(tenderSystem.projectById[projectId], userId)
	tenderSystem.projectByPrice[projectId] = append(tenderSystem.projectByPrice[projectId], price)
	fmt.Println("add", projectId, userId, tenderSystem.projectById, projectId, price, tenderSystem.projectByPrice)
}

func (tenderSystem *TenderSystem) updateTender(userId int, projectId int, price int) int {
	userList := tenderSystem.projectById[projectId]
	fmt.Println(userList)
	priceList := tenderSystem.projectByPrice[projectId]
	for i := range userList {
		if userList[i] == userId {
			prevPrice := priceList[i]
			if len(priceList) == 0 {
				fmt.Println("1")
				priceList[i] = price
			} else {
				fmt.Println("2")
				tenderSystem.projectById[projectId] = append(tenderSystem.projectById[projectId][:i], tenderSystem.projectById[projectId][i+1:]...)
				tenderSystem.projectByPrice[projectId] = append(tenderSystem.projectByPrice[projectId][:i], tenderSystem.projectByPrice[projectId][i+1:]...)
				tenderSystem.projectById[projectId] = append(tenderSystem.projectById[projectId], userId)
				tenderSystem.projectByPrice[projectId] = append(tenderSystem.projectByPrice[projectId], price)
			}
			fmt.Println("updata", priceList, tenderSystem.projectByPrice[projectId])
			return prevPrice
		}
	}
	return -1
}

func (tenderSystem *TenderSystem) removeTender(userId int, projectId int) int {
	userList := tenderSystem.projectById[projectId]
	priceList := tenderSystem.projectByPrice[projectId]
	fmt.Println("REMOVE", userList, priceList)
	for i := range userList {
		if userList[i] == userId {
			prevPrice := priceList[i]
			userList = append(userList[:i], userList[i+1:]...)
			priceList = append(priceList[:i], priceList[i+1:]...)
			return prevPrice
		}
	}
	return -1
}

func (tenderSystem *TenderSystem) queryTender(projectId int, price int) int {
	userList := tenderSystem.projectById[projectId]
	priceList := tenderSystem.projectByPrice[projectId]
	index := -1
	fmt.Println(priceList)
	for i := range priceList {
		if priceList[i] > price {
			if index == -1 {
				index = i
			} else if priceList[i] < priceList[index] {
				index = i
			}
		}
	}
	if index == -1 {
		return -1
	}
	return userList[index]
}
