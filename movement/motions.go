package movement

import "fmt"

const homePositionCoord = 0.00

func HomePosition() ([2]float32, string) {
	homingPositionXY := [2]float32{homePositionCoord, homePositionCoord}
	homingPositionG := fmt.Sprintf("G00 X%.2f Y%.2f", homingPositionXY[0], homingPositionXY[1])

	return homingPositionXY, homingPositionG

}

func StartMovementFullSpeed(origin [2]float32, distanceX float32, distanceY float32) ([2]float32, string) {
	nevPositionXY := [2]float32{origin[0] + distanceX, origin[1] + distanceY}
	newPositionG := fmt.Sprintf("G00 X%.2f Y%.2f", nevPositionXY[0], nevPositionXY[1])

	return nevPositionXY, newPositionG
}
func StartMovementVariableSpeed(origin [2]float32, distanceX float32, distanceY float32, speed float32) ([2]float32, string) {
	nevPositionXY := [2]float32{origin[0] + distanceX, origin[1] + distanceY}
	newPositionG := fmt.Sprintf("G01 X%.2f Y%.2f F%.2f", nevPositionXY[0], nevPositionXY[1], speed)

	return nevPositionXY, newPositionG
}
func ContinueMovmentX(origin [2]float32, distanceX float32) ([2]float32, string) {

	nevPositionXY := [2]float32{origin[0] + distanceX, origin[1]}
	newPositionG := fmt.Sprintf("X%.2f", nevPositionXY[0])

	return nevPositionXY, newPositionG
}
func ContinueMovmentY(origin [2]float32, distanceY float32) ([2]float32, string) {

	nevPositionXY := [2]float32{origin[0], origin[1] + distanceY}
	newPositionG := fmt.Sprintf("Y%.2f", nevPositionXY[1])

	return nevPositionXY, newPositionG
}
