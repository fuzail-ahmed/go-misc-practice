package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	//panic("Please implement the CanFastAttack() function")
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	//panic("Please implement the CanSpy() function")
    if knightIsAwake || archerIsAwake || prisonerIsAwake {
        return true
    } else {
        return false
    }
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	//panic("Please implement the CanSignalPrisoner() function")
    if archerIsAwake == false && prisonerIsAwake == true {
        return true
    } else {
        return false
    }
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {

    if prisonerIsAwake == true && knightIsAwake == false && archerIsAwake == false {
        return true
    } else if prisonerIsAwake == false && knightIsAwake == true && archerIsAwake == false {
        if petDogIsPresent {
            return true
        } else {
        	return false
        }
    } else if prisonerIsAwake == false && knightIsAwake == false && archerIsAwake == false {
        if petDogIsPresent {
            return true
        } else {
        	return false
        }
    } else if prisonerIsAwake == true && knightIsAwake == true && archerIsAwake == false {
        if petDogIsPresent {
            return true
        } else {
        	return false
        }
    } else {
        return false
    }
}
