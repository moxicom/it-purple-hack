package trees

var categoryID int64

func GetCategoryTree() *Node {
	// Create root category
	categoryID = 0
	allRegions := NewNode(rawCategories.Name, categoryID)
	for _, subCategory := range rawCategories.SubElements {
		allRegions.insertNodeRecursively(subCategory, &categoryID)
	}

	return allRegions
}

var rawCategories = RawData{
	Name: "Категории товаров",
	SubElements: []RawData{
		{
			Name: "Бытовая электроника",
			SubElements: []RawData{
				{Name: "Товары для компьютера"},
				{Name: "Фототехника"},
				{Name: "Телефоны"},
				{Name: "Планшеты и электронные книги"},
				{Name: "Оргтехника и расходники"},
				{Name: "Ноутбуки"},
				{Name: "Настольные компьютеры"},
				{Name: "Игры, приставки и программы"},
				{Name: "Аудио и видео"},
			},
		},
		{
			Name: "Готовый бизнес и оборудование",
			SubElements: []RawData{
				{Name: "Готовый бизнес"},
				{Name: "Оборудование для бизнеса"},
			},
		},
		{
			Name: "Для дома и дачи",
			SubElements: []RawData{
				{Name: "Мебель и интерьер"},
				{Name: "Ремонт и строительство"},
				{Name: "Продукты питания"},
				{Name: "Растения"},
				{Name: "Бытовая техника"},
				{Name: "Посуда и товары для кухни"},
			},
		},
		// Другие категории здесь...
	},
}
