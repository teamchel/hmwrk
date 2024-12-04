package main

import (
	"fmt"
	"time"
)

const (
	MInKm      = 1000
	MinInHours = 60
	LenStep    = 0.65
	CmInM      = 100
)

type Training struct {
	TrainingType string
	Action       int
	LenStep      float64
	Duration     time.Duration
	Weight       int
}

func (t Training) distance() float64 {

	return float64(t.Action) * LenStep / float64(MInKm)
}

func (t Training) meanSpeed() float64 {

	d := t.distance()
	return d / float64(t.Duration)
}

func (t Training) Calories() float64 {

	return 0
}

type InfoMessage struct {
	TrainingType string
	Duration     time.Duration
	Distance     float64
	Speed        float64
	Calories     float64
}

func (t Training) TrainingInfo() InfoMessage {
	// вставьте ваш код ниже
	return InfoMessage{
		t.TrainingType,
		t.Duration,
		t.distance(),
		t.meanSpeed(),
		t.Calories(),
	}
}

func (i InfoMessage) String() string {
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %v мин\nДистанция: %.2f км.\nСр. скорость: %.2f км/ч\nПотрачено ккал: %.2f\n",
		i.TrainingType,
		i.Duration.Minutes(),
		i.Distance,
		i.Speed,
		i.Calories,
	)
}

type CaloriesCalculator interface {
	Calories() float64
	TrainingInfo() InfoMessage
}

const (
	CaloriesMeanSpeedMultiplier = 18
	CaloriesMeanSpeedShift      = 1.79
)

type Running struct {
	Training
}

func (r Running) Calories() float64 {

	return ((CaloriesMeanSpeedMultiplier*r.meanSpeed() + CaloriesMeanSpeedShift) * float64(r.Weight) / MInKm * float64(r.Duration) * MinInHours)
}

func (r Running) TrainingInfo() InfoMessage {

	return InfoMessage{
		r.TrainingType,
		r.Duration,
		r.distance(),
		r.meanSpeed(),
		r.Calories(),
	}
}

const (
	CaloriesWeightMultiplier      = 0.035
	CaloriesSpeedHeightMultiplier = 0.029
	KmHInMsec                     = 0.278
)

type Walking struct {
	Training
	Height float64
	Weight float64
}

func (w Walking) Calories() float64 {

	return ((CaloriesWeightMultiplier*w.Weight + (w.meanSpeed()*w.meanSpeed()/w.Height)*CaloriesSpeedHeightMultiplier*w.Weight) * float64(w.Duration) * MinInHours)

}

func (w Walking) TrainingInfo() InfoMessage {

	return InfoMessage{
		w.TrainingType,
		w.Duration,
		w.distance(),
		w.meanSpeed(),
		w.Calories(),
	}
}

const (
	SwimmingLenStep                  = 1.38
	SwimmingCaloriesMeanSpeedShift   = 1.1
	SwimmingCaloriesWeightMultiplier = 2
)

type Swimming struct {
	Training
	LengthPool int
	CountPool  int
}

func (s Swimming) meanSpeed() float64 {

	mSpeed := float64(s.LengthPool) * float64(s.CountPool) / float64(MInKm) / float64(s.Duration)
	return mSpeed
}

func (s Swimming) Calories() float64 {
	calories := (s.meanSpeed() + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * float64(s.Weight) * float64(s.Duration)
	return calories
}

func (s Swimming) TrainingInfo() InfoMessage {

	return InfoMessage{
		s.TrainingType,
		s.Duration,
		s.distance(),
		s.meanSpeed(),
		s.Calories(),
	}
}

func ReadData(training CaloriesCalculator) string {

	calories := training.Calories()

	info := training.TrainingInfo()

	info.Calories = calories
	return fmt.Sprint(info)
}

func main() {

	swimming := Swimming{
		Training: Training{
			TrainingType: "Плавание",
			Action:       2000,
			LenStep:      SwimmingLenStep,
			Duration:     90 * time.Minute,
			Weight:       85,
		},
		LengthPool: 50,
		CountPool:  5,
	}

	fmt.Println(ReadData(swimming))

	walking := Walking{
		Training: Training{
			TrainingType: "Ходьба",
			Action:       20000,
			LenStep:      LenStep,
			Duration:     3*time.Hour + 45*time.Minute,
			Weight:       85,
		},
		Height: 185,
	}

	fmt.Println(ReadData(walking))

	running := Running{
		Training: Training{
			TrainingType: "Бег",
			Action:       5000,
			LenStep:      LenStep,
			Duration:     30 * time.Minute,
			Weight:       85,
		},
	}

	fmt.Println(ReadData(running))

}
