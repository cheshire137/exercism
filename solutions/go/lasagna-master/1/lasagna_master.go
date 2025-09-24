package lasagna

func PreparationTime(layers []string, prepTimePerLayerInMin int) int {
    if prepTimePerLayerInMin < 1 {
        prepTimePerLayerInMin = 2
    }
    return len(layers) * prepTimePerLayerInMin
}

func Quantities(layers []string) (int, float64) {
    noodleLayers := 0
    sauceLayers := 0
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodleLayers++
        } else if layers[i] == "sauce" {
            sauceLayers++
        }
    }
    return noodleLayers * 50, float64(sauceLayers) * 0.2
}

func AddSecretIngredient(friendIngredients []string, myIngredients []string) {
    fromIdx := len(friendIngredients) - 1
    toIdx := len(myIngredients) - 1
    myIngredients[toIdx] = friendIngredients[fromIdx]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaledQuantities := make([]float64, len(quantities))
    factor := float64(portions) / float64(2)
    for i := 0; i < len(quantities); i++ {
        scaledQuantities[i] = quantities[i] * factor
    }
    return scaledQuantities
}
