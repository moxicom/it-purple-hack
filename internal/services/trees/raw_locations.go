package trees

var rawLocations = RawData{
	Name: "Россия",
	SubElements: []RawData{
		{
			Name: "Московская область",
			SubElements: []RawData{
				{
					Name: "Москва",
					SubElements: []RawData{
						{
							Name: "Арбат",
							SubElements: []RawData{
								{Name: "улица Арбат"},
								{Name: "улица Старый Арбат"},
							},
						},
						{
							Name: "Пресненский район",
							SubElements: []RawData{
								{Name: "улица Пресненский Вал"},
								{Name: "улица Баррикадная"},
							},
						},
					},
				},
				{
					Name: "Подмосковье",
					SubElements: []RawData{
						{
							Name: "Красногорск",
							SubElements: []RawData{
								{Name: "улица Ленина"},
								{Name: "улица Гагарина"},
							},
						},
						{
							Name: "Химки",
							SubElements: []RawData{
								{Name: "улица Центральная"},
								{Name: "улица Школьная"},
							},
						},
					},
				},
			},
		},
		{
			Name: "Волгоградская область",
			SubElements: []RawData{
				{
					Name: "Калифорния",
					SubElements: []RawData{
						{
							Name: "Лос-Анджелес",
							SubElements: []RawData{
								{
									Name: "Голливуд",
									SubElements: []RawData{
										{Name: "Sunset Boulevard"},
										{Name: "Hollywood Boulevard"},
									},
								},
							},
						},
						{
							Name: "Сан-Франциско",
							SubElements: []RawData{
								{
									Name: "Даунтаун",
									SubElements: []RawData{
										{Name: "Market Street"},
										{Name: "Mission Street"},
									},
								},
								{
									Name: "Силликоновая Долина",
									SubElements: []RawData{
										{Name: "El Camino Real"},
										{Name: "California Avenue"},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
