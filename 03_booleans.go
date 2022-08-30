package main

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	var canFastAttack bool = !knightIsAwake
	return canFastAttack
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	var canSpy bool = knightIsAwake || archerIsAwake || prisonerIsAwake
	return canSpy
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	var canSignalPrisoner bool = !archerIsAwake && prisonerIsAwake
	return canSignalPrisoner
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	var canFreePrisonder bool = ((prisonerIsAwake && !archerIsAwake && !knightIsAwake) || (petDogIsPresent && !archerIsAwake))
	return canFreePrisonder
}
