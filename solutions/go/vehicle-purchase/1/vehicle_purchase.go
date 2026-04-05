package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// panic("NeedsLicense not implemented")
    return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	// panic("ChooseVehicle not implemented")
    var res_opt string
    if option1 < option2 {
        res_opt = option1
    } else {
        res_opt = option2
    }
    return res_opt + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// panic("CalculateResellPrice not implemented")
    if age < 3 {
        return originalPrice * 0.8
    } else if age < 10 {
        return originalPrice * 0.7
    } else {
        return originalPrice * 0.5
    }
}
