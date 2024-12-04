package main

import (
	"fmt"
	"time"
)

// Общие константы для вычислений.
const (
	MInKm      = 1000 // количество метров в одном километре
	MinInHours = 60   // количество минут в одном часе
	LenStep    = 0.65 // длина одного шага
	CmInM      = 100  // количество сантиметров в одном метре
)

// Training общая структура для всех тренировок
type Training struct {
	TrainingType string
	Action       int
	LenStep      float64
	Duration     time.Duration
	Weight       int
}

// distance возвращает дистанцию, которую преодолел пользователь.
// Формула расчета:
// количество_повторов * длина_шага / м_в_км
func (t Training) distance() float64 {
	// вставьте ваш код ниже

	return float64(t.Action) * LenStep / float64(MInKm)
}

// meanSpeed возвращает среднюю скорость бега или ходьбы.
func (t Training) meanSpeed() float64 {
	// вставьте ваш код ниже
	d := t.distance()
	return d / float64(t.Duration)
}

// Calories возвращает количество потраченных килокалорий на тренировке.
// Пока возвращаем 0, так как этот метод будет переопределяться для каждого типа тренировки.
func (t Training) Calories() float64 {
	// вставьте ваш код ниже
	return 0
}

// InfoMessage содержит информацию о проведенной тренировке.
type InfoMessage struct {
	// добавьте необходимые поля в структуру
	TrainingType string        // тип тренировки
	Duration     time.Duration // длительность тренировки
	Distance     float64       // расстояние, которое преодолел пользователь
	Speed        float64       // средняя скорость, с которой двигался пользователь
	Calories     float64       // количество потраченных килокалорий на тренировке
}

// TrainingInfo возвращает труктуру InfoMessage, в которой хранится вся информация о проведенной тренировке.
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

// String возвращает строку с информацией о проведенной тренировке.
func (i InfoMessage) String() string {
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %v мин\nДистанция: %.2f км.\nСр. скорость: %.2f км/ч\nПотрачено ккал: %.2f\n",
		i.TrainingType,
		i.Duration.Minutes(),
		i.Distance,
		i.Speed,
		i.Calories,
	)
}

// CaloriesCalculator интерфейс для структур: Running, Walking и Swimming.
type CaloriesCalculator interface {
	// добавьте необходимые методы в интерфейс
	Calories() float64
	TrainingInfo() InfoMessage
}

// Константы для расчета потраченных килокалорий при беге.
const (
	CaloriesMeanSpeedMultiplier = 18   // множитель средней скорости бега
	CaloriesMeanSpeedShift      = 1.79 // коэффициент изменения средней скорости
)

// Running структура, описывающая тренировку Бег.
type Running struct {
	// добавьте необходимые поля в структуру
	Training
}

// Calories возввращает количество потраченных килокалория при беге.
// Формула расчета:
// ((18 * средняя_скорость_в_км/ч + 1.79) * вес_спортсмена_в_кг / м_в_км * время_тренировки_в_часах * мин_в_часе)
// Это переопределенный метод Calories() из Training.
func (r Running) Calories() float64 {
	// вставьте ваш код ниже

	return ((CaloriesMeanSpeedMultiplier*r.meanSpeed() + CaloriesMeanSpeedShift) * float64(r.Weight) / MInKm * float64(r.Duration) * MinInHours)
}

// TrainingInfo возвращает структуру InfoMessage с информацией о проведенной тренировке.
// Это переопределенный метод TrainingInfo() из Training.
func (r Running) TrainingInfo() InfoMessage {
	// вставьте ваш код ниже
	return InfoMessage{
		r.TrainingType,
		r.Duration,
		r.distance(),
		r.meanSpeed(),
		r.Calories(),
	}
}

// Константы для расчета потраченных килокалорий при ходьбе.
const (
	CaloriesWeightMultiplier      = 0.035 // коэффициент для веса
	CaloriesSpeedHeightMultiplier = 0.029 // коэффициент для роста
	KmHInMsec                     = 0.278 // коэффициент для перевода км/ч в м/с
)

// Walking структура описывающая тренировку Ходьба
type Walking struct {
	// добавьте необходимые поля в структуру
	Training
	Height float64 // рост пользователя
	Weight float64 // вес пользователя
}

// Calories возвращает количество потраченных килокалорий при ходьбе.
// Формула расчета:
// ((0.035 * вес_спортсмена_в_кг + (средняя_скорость_в_метрах_в_секунду**2 / рост_в_метрах)
// * 0.029 * вес_спортсмена_в_кг) * время_тренировки_в_часах * мин_в_ч)
// Это переопределенный метод Calories() из Training.
func (w Walking) Calories() float64 {
	// вставьте ваш код ниже

	return ((CaloriesWeightMultiplier*w.Weight + (w.meanSpeed()*w.meanSpeed()/w.Height)*CaloriesSpeedHeightMultiplier*w.Weight) * float64(w.Duration) * MinInHours)

}

// TrainingInfo возвращает структуру InfoMessage с информацией о проведенной тренировке.
// Это переопределенный метод TrainingInfo() из Training.
func (w Walking) TrainingInfo() InfoMessage {
	// вставьте ваш код ниже
	return InfoMessage{
		w.TrainingType,
		w.Duration,
		w.distance(),
		w.meanSpeed(),
		w.Calories(),
	}
}

// Константы для расчета потраченных килокалорий при плавании.
const (
	SwimmingLenStep                  = 1.38 // длина одного гребка
	SwimmingCaloriesMeanSpeedShift   = 1.1  // коэффициент изменения средней скорости
	SwimmingCaloriesWeightMultiplier = 2    // множитель веса пользователя
)

// Swimming структура, описывающая тренировку Плавание
type Swimming struct {
	// добавьте необходимые поля в структуру
	Training
	LengthPool int // длина бассейна
	CountPool  int // количество пересечений бассейна
}

// meanSpeed возвращает среднюю скорость при плавании.
// Формула расчета:
// длина_бассейна * количество_пересечений / м_в_км / продолжительность_тренировки
// Это переопределенный метод Calories() из Training.
func (s Swimming) meanSpeed() float64 {
	// вставьте ваш код ниже
	mSpeed := float64(s.LengthPool) * float64(s.CountPool) / float64(MInKm) / float64(s.Duration)
	return mSpeed
}

// Calories возвращает количество калорий, потраченных при плавании.
// Формула расчета:
// (средняя_скорость_в_км/ч + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * вес_спортсмена_в_кг * время_тренировки_в_часах
// Это переопределенный метод Calories() из Training.
func (s Swimming) Calories() float64 {
	// вставьте ваш код ниже
	calories := (s.meanSpeed() + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * float64(s.Weight) * float64(s.Duration)
	return calories
}

// TrainingInfo returns info about swimming training.
// Это переопределенный метод TrainingInfo() из Training.
func (s Swimming) TrainingInfo() InfoMessage {
	// вставьте ваш код ниже
	return InfoMessage{
		s.TrainingType,
		s.Duration,
		s.distance(),
		s.meanSpeed(),
		s.Calories(),
	}
}

// ReadData возвращает информацию о проведенной тренировке.
// ReadData возвращает информацию о проведенной тренировке.
func ReadData(training CaloriesCalculator) string {
	// получите количество затраченных калорий
	calories := training.Calories()
	// получите информацию о тренировке
	info := training.TrainingInfo()
	// добавьте полученные калории в структуру с информацией о тренировке
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
