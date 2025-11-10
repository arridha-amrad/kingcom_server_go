package seedcmd

import (
	"kingcom_api/internal/dto"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/models"
	"kingcom_api/internal/utils"
	"math"

	"github.com/brianvoe/gofakeit/v6"

	"golang.org/x/crypto/bcrypt"
)

func Run() {
	logger := lib.GetLogger()
	db := initDB(logger)
	seedUsers(db, logger)
	seedProducts(db, logger)
}

func seedUsers(db *lib.Database, logger *lib.Logger) {
	total := 20
	users := make([]models.User, 0, total)
	for range total {
		pwd, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
		u := models.User{
			Username:   gofakeit.Username(),
			Name:       gofakeit.Name(),
			Email:      gofakeit.Email(),
			Password:   string(pwd),
			Provider:   "credentials",
			Role:       "user",
			JwtVersion: gofakeit.UUID(),
			IsVerified: true,
		}
		users = append(users, u)
	}
	if err := db.Create(&users).Error; err != nil {
		logger.Panic(err)
	}
	logger.Info("users seeeded")
}

func initDB(logger *lib.Logger) *lib.Database {
	env := lib.NewEnv()
	db := lib.NewDatabase(env, logger)
	return db
}

func seedProducts(db *lib.Database, logger *lib.Logger) {
	// products := []dto.CreateProduct{
	// 	{
	// 		Name:          "Logitech G Pro X Superlight Wireless Mouse",
	// 		Weight:        0.063,
	// 		Price:         1890000,
	// 		Description:   "Mouse gaming nirkabel super ringan dengan sensor HERO 25K dan desain minimalis untuk akurasi tinggi.",
	// 		Stock:         45,
	// 		Discount:      10,
	// 		Specification: ptr("Sensor: HERO 25K, DPI: 100–25.600, Koneksi: Wireless LIGHTSPEED, Baterai: 70 jam"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=tN7iNn1QX7E"),
	// 		Images: []string{
	// 			"https://example.com/images/logitech-gpro-x.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "Razer BlackWidow V4 Pro Mechanical Keyboard",
	// 		Weight:        1.15,
	// 		Price:         3999000,
	// 		Description:   "Keyboard mekanik RGB dengan switch Razer Green, makro khusus, dan pergelangan tangan magnetik.",
	// 		Stock:         30,
	// 		Discount:      5,
	// 		Specification: ptr("Switch: Razer Green, Backlight: RGB, Connectivity: USB-C detachable, N-key rollover"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=YkZmFqkDjBw"),
	// 		Images: []string{
	// 			"https://example.com/images/razer-blackwidow-v4-pro.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "ASUS ROG Strix B550-F Gaming Motherboard",
	// 		Weight:        1.8,
	// 		Price:         2899000,
	// 		Description:   "Motherboard AM4 dengan dukungan PCIe 4.0 dan fitur pendinginan premium untuk performa gaming stabil.",
	// 		Stock:         20,
	// 		Discount:      0,
	// 		Specification: ptr("Chipset: AMD B550, Socket: AM4, RAM: 128GB DDR4, PCIe 4.0 x16"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=Qw48cE9zW2A"),
	// 		Images: []string{
	// 			"https://example.com/images/asus-rog-strix-b550f.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "AMD Ryzen 7 7800X3D Processor",
	// 		Weight:        0.07,
	// 		Price:         6499000,
	// 		Description:   "Prosesor 8-core dengan 3D V-Cache untuk performa gaming luar biasa dan efisiensi daya tinggi.",
	// 		Stock:         35,
	// 		Discount:      7,
	// 		Specification: ptr("Cores: 8, Threads: 16, Base Clock: 4.2GHz, Boost: 5.0GHz, Cache: 104MB"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=cjNfV8ZzQGQ"),
	// 		Images: []string{
	// 			"https://example.com/images/ryzen-7-7800x3d.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "NVIDIA GeForce RTX 4070 Ti Super",
	// 		Weight:        1.2,
	// 		Price:         13999000,
	// 		Description:   "GPU high-end untuk gaming 4K dengan DLSS 3, ray tracing, dan pendinginan triple-fan.",
	// 		Stock:         15,
	// 		Discount:      8,
	// 		Specification: ptr("VRAM: 16GB GDDR6X, CUDA Cores: 7680, Clock: 2.6GHz Boost, Power: 285W"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=8PrHjEEl6xE"),
	// 		Images: []string{
	// 			"https://example.com/images/rtx-4070ti-super.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "Corsair Vengeance RGB 32GB DDR5 6000MHz",
	// 		Weight:        0.12,
	// 		Price:         2399000,
	// 		Description:   "RAM DDR5 berkecepatan tinggi dengan pencahayaan RGB dan kompatibilitas XMP 3.0.",
	// 		Stock:         60,
	// 		Discount:      10,
	// 		Specification: ptr("Capacity: 32GB (2x16GB), Speed: 6000MHz, Type: DDR5, CAS Latency: 36"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=Bx_gSuEJw_4"),
	// 		Images: []string{
	// 			"https://example.com/images/corsair-vengeance-rgb-ddr5.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "Samsung 990 PRO 1TB NVMe SSD",
	// 		Weight:        0.009,
	// 		Price:         2599000,
	// 		Description:   "SSD NVMe PCIe 4.0 dengan kecepatan baca hingga 7450MB/s untuk loading super cepat.",
	// 		Stock:         50,
	// 		Discount:      5,
	// 		Specification: ptr("Interface: PCIe 4.0 x4, Read: 7450MB/s, Write: 6900MB/s, TBW: 600"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=Z4jdrK2g6jU"),
	// 		Images: []string{
	// 			"https://example.com/images/samsung-990pro.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "Cooler Master Hyper 212 Halo Black",
	// 		Weight:        0.65,
	// 		Price:         599000,
	// 		Description:   "Pendingin CPU legendaris dengan desain baru dan kipas ARGB 120mm.",
	// 		Stock:         70,
	// 		Discount:      0,
	// 		Specification: ptr("Fan Speed: 2000 RPM, Noise: 27dBA, Compatibility: Intel & AMD sockets"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=Y2KMpCXvArE"),
	// 		Images: []string{
	// 			"https://example.com/images/hyper-212-halo.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "NZXT H7 Flow Mid Tower Case",
	// 		Weight:        8.2,
	// 		Price:         1899000,
	// 		Description:   "Casing PC dengan airflow optimal dan desain minimalis khas NZXT.",
	// 		Stock:         25,
	// 		Discount:      0,
	// 		Specification: ptr("Material: Steel, Tempered Glass, ATX compatible, 2 pre-installed fans"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=FxI6C24sW-Y"),
	// 		Images: []string{
	// 			"https://example.com/images/nzxt-h7-flow.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "MSI MAG A850GL PCIE5 Power Supply",
	// 		Weight:        2.4,
	// 		Price:         1699000,
	// 		Description:   "PSU 850W 80+ Gold dengan konektor PCIe 5.0 dan kabel modular penuh.",
	// 		Stock:         40,
	// 		Discount:      5,
	// 		Specification: ptr("Wattage: 850W, Efficiency: 80+ Gold, Type: Fully Modular, Connector: PCIe 5.0 16-pin"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=Cr2HnDlpRHQ"),
	// 		Images: []string{
	// 			"https://example.com/images/msi-mag-a850gl.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "SteelSeries Arctis Nova 7 Wireless Headset",
	// 		Weight:        0.35,
	// 		Price:         2899000,
	// 		Description:   "Headset gaming wireless dengan 2.4GHz + Bluetooth ganda dan kualitas audio premium.",
	// 		Stock:         32,
	// 		Discount:      10,
	// 		Specification: ptr("Drivers: 40mm Neodymium, Battery: 38 jam, Mic: Retractable ClearCast Gen2"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=0Wgz-Sk6Quk"),
	// 		Images: []string{
	// 			"https://example.com/images/arctis-nova7.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "Elgato Stream Deck MK.2",
	// 		Weight:        0.3,
	// 		Price:         2999000,
	// 		Description:   "Kontroler streaming dengan 15 tombol LCD yang bisa diprogram untuk shortcut dan makro.",
	// 		Stock:         18,
	// 		Discount:      0,
	// 		Specification: ptr("Keys: 15 LCD, Connection: USB-C, Custom Profiles: Yes, Software: Stream Deck App"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=YjOUpPYzv64"),
	// 		Images: []string{
	// 			"https://example.com/images/elgato-streamdeck.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "AOC 27G2SPU 27\" 165Hz Gaming Monitor",
	// 		Weight:        5.1,
	// 		Price:         2799000,
	// 		Description:   "Monitor gaming Full HD 165Hz dengan panel IPS dan waktu respons 1ms.",
	// 		Stock:         22,
	// 		Discount:      12,
	// 		Specification: ptr("Panel: IPS, Refresh: 165Hz, Response: 1ms, Ports: HDMI/DP"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=hZkZ6J7XcX0"),
	// 		Images: []string{
	// 			"https://example.com/images/aoc-27g2spu.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "ASUS TUF Gaming VG27AQ 27\" 2K Monitor",
	// 		Weight:        6.1,
	// 		Price:         5699000,
	// 		Description:   "Monitor 2K 165Hz dengan teknologi ELMB Sync untuk pengalaman gaming bebas blur.",
	// 		Stock:         10,
	// 		Discount:      5,
	// 		Specification: ptr("Resolution: 2560x1440, Refresh: 165Hz, HDR10, Adaptive-Sync"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=vQoMi5R9mLE"),
	// 		Images: []string{
	// 			"https://example.com/images/asus-vg27aq.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "Kingston FURY Renegade 2TB NVMe SSD",
	// 		Weight:        0.008,
	// 		Price:         3999000,
	// 		Description:   "SSD gaming performa tinggi untuk sistem ekstrem dengan kecepatan baca hingga 7300MB/s.",
	// 		Stock:         25,
	// 		Discount:      7,
	// 		Specification: ptr("Interface: PCIe 4.0, Read: 7300MB/s, Write: 7000MB/s, TBW: 1600"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=sh6ZBqpKkDs"),
	// 		Images: []string{
	// 			"https://example.com/images/kingston-renegade-2tb.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "HyperX Cloud III Wireless Gaming Headset",
	// 		Weight:        0.34,
	// 		Price:         2799000,
	// 		Description:   "Headset gaming nirkabel dengan DTS Headphone:X dan daya tahan baterai 120 jam.",
	// 		Stock:         37,
	// 		Discount:      8,
	// 		Specification: ptr("Driver: 53mm, Battery: 120 jam, Connectivity: 2.4GHz Wireless"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=afg4V_5D5eE"),
	// 		Images: []string{
	// 			"https://example.com/images/hyperx-cloud3.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "Logitech C920 HD Pro Webcam",
	// 		Weight:        0.162,
	// 		Price:         1099000,
	// 		Description:   "Webcam Full HD 1080p dengan autofokus dan mikrofon stereo bawaan.",
	// 		Stock:         50,
	// 		Discount:      15,
	// 		Specification: ptr("Resolution: 1080p30, Connection: USB-A, Field of View: 78°"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=f7cQoMtzEec"),
	// 		Images: []string{
	// 			"https://example.com/images/logitech-c920.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "Gigabyte AORUS M5 RGB Gaming Mouse",
	// 		Weight:        0.13,
	// 		Price:         899000,
	// 		Description:   "Mouse gaming ergonomis dengan sensor 16000 DPI dan pencahayaan RGB penuh.",
	// 		Stock:         55,
	// 		Discount:      12,
	// 		Specification: ptr("Sensor: PWM3389, DPI: 16000, Switch: Omron, Weight Adjust: Yes"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=By5jRvGk4u8"),
	// 		Images: []string{
	// 			"https://example.com/images/aorus-m5.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "Seagate Barracuda 2TB 7200RPM HDD",
	// 		Weight:        0.45,
	// 		Price:         949000,
	// 		Description:   "Hard disk desktop 2TB dengan kecepatan 7200RPM dan buffer 256MB.",
	// 		Stock:         70,
	// 		Discount:      5,
	// 		Specification: ptr("Capacity: 2TB, RPM: 7200, Cache: 256MB, Interface: SATA 6Gb/s"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=0sZsmi87gDk"),
	// 		Images: []string{
	// 			"https://example.com/images/seagate-barracuda-2tb.jpg",
	// 		},
	// 	},
	// 	{
	// 		Name:          "DeepCool AK400 CPU Cooler",
	// 		Weight:        0.65,
	// 		Price:         499000,
	// 		Description:   "Pendingin CPU tower dengan performa tinggi dan desain kompak.",
	// 		Stock:         90,
	// 		Discount:      0,
	// 		Specification: ptr("Fan: 120mm PWM, Noise: 29dBA, TDP: 220W"),
	// 		VideoUrl:      ptr("https://www.youtube.com/watch?v=zyaxZp8ZT4I"),
	// 		Images: []string{
	// 			"https://example.com/images/deepcool-ak400.jpg",
	// 		},
	// 	},
	// }

	totalProducts := 50
	products := make([]dto.CreateProduct, 0, totalProducts)
	for range totalProducts {
		product := dto.CreateProduct{
			Name:          gofakeit.ProductName(),
			Weight:        math.Ceil(gofakeit.Float64Range(0.5, 3.0)*10) / 10,
			Price:         float64(gofakeit.Number(1500, 100000)),
			Description:   gofakeit.ProductDescription(),
			Stock:         gofakeit.UintRange(5, 100),
			Discount:      gofakeit.IntRange(5, 30),
			Specification: ptr(gofakeit.ProductFeature()),
			Images:        []string{"https://www.differencebetween.net/wp-content/uploads/2012/01/Difference-Between-Example-and-Sample.jpg"},
		}
		products = append(products, product)
	}
	productsDomain := make([]models.Product, 0, len(products))
	for _, p := range products {
		slug := utils.ToSlug(p.Name)
		totalImage := 2
		images := make([]models.ProductImage, 0, totalImage)
		for range totalImage {
			images = append(images, models.ProductImage{
				Url: "https://www.differencebetween.net/wp-content/uploads/2012/01/Difference-Between-Example-and-Sample.jpg",
			})
		}
		productsDomain = append(productsDomain, models.Product{
			Name:          p.Name,
			Weight:        p.Weight,
			Slug:          slug,
			Price:         p.Price,
			Description:   p.Description,
			Stock:         p.Stock,
			Discount:      p.Discount,
			Specification: p.Specification,
			VideoUrl:      p.VideoUrl,
			Images:        images,
		})
	}
	if err := db.Create(&productsDomain).Error; err != nil {
		logger.Panic(err)
	}
	logger.Info("products seeded")
}

func ptr(s string) *string { return &s }
