package fictions

import (
	"math/rand"
	"regexp"
	"strings"
)

type Type string

const (
	TypeAnime  Type = "anime"
	TypeMovie  Type = "movie"
	TypeSeries Type = "series"
)

type Fiction struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
	Slug string `json:"slug"`
}

var List = []Fiction{
	{Name: "Наруто", Type: TypeAnime, Slug: "naruto"},
	{Name: "Вторжение титанов", Type: TypeAnime, Slug: "attack-on-titan"},
	{Name: "Моя геройская академия", Type: TypeAnime, Slug: "my-hero-academia"},
	{Name: "Тетрадь смерти", Type: TypeAnime, Slug: "death-note"},
	{Name: "Истребитель демонов", Type: TypeAnime, Slug: "demon-slayer"},
	{Name: "Блич", Type: TypeAnime, Slug: "bleach"},
	{Name: "Охотник × Охотник", Type: TypeAnime, Slug: "hunter-x-hunter"},
	{Name: "Судьба/Остаться ночью: Безграничные клинки", Type: TypeAnime, Slug: "fate-stay-night-unlimited-blade-works"},
	{Name: "Магическая битва", Type: TypeAnime, Slug: "jujutsu-kaisen"},
	{Name: "Сага о Винланде", Type: TypeAnime, Slug: "vinland-saga"},
	{Name: "Титаник", Type: TypeMovie, Slug: "titanic"},
	{Name: "Я легенда", Type: TypeMovie, Slug: "i-am-a-legend"},
	{Name: "Матрица", Type: TypeMovie, Slug: "the-matrix"},
	{Name: "Интерстеллар", Type: TypeMovie, Slug: "interstellar"},
	{Name: "Аватар", Type: TypeMovie, Slug: "avatar"},
	{Name: "Стражи Галактики", Type: TypeMovie, Slug: "guardians-of-the-galaxy"},
	{Name: "Одержимость", Type: TypeMovie, Slug: "whiplash"},
	{Name: "Железный человек", Type: TypeMovie, Slug: "iron-man"},
	{Name: "Человек-паук", Type: TypeMovie, Slug: "spider-man"},
	{Name: "Мстители: Война бесконечности", Type: TypeMovie, Slug: "avengers-infinity-war"},
	{Name: "Мстители: Финал", Type: TypeMovie, Slug: "avengers-endgame"},
	{Name: "Стражи Галактики 2", Type: TypeMovie, Slug: "guardians-of-the-galaxy-vol-2"},
	{Name: "Бумажный дом", Type: TypeSeries, Slug: "money-heist"},
	{Name: "Во все тяжкие", Type: TypeSeries, Slug: "breaking-bad"},
	{Name: "Очень странные дела", Type: TypeSeries, Slug: "stranger-things"},
	{Name: "Клан Сопрано", Type: TypeSeries, Slug: "the-sopranos"},
	{Name: "Острые козырьки", Type: TypeSeries, Slug: "peaky-blinders"},
	{Name: "Побег", Type: TypeSeries, Slug: "prison-break"},
	{Name: "Офис", Type: TypeSeries, Slug: "the-office"},
	{Name: "Шерлок", Type: TypeSeries, Slug: "sherlock"},
	{Name: "Декстер", Type: TypeSeries, Slug: "dexter"},
	{Name: "Пацаны", Type: TypeSeries, Slug: "the-boys"},
}

var slugRe = regexp.MustCompile(`[^a-z0-9-]+`)

func Slug(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "-")
	s = slugRe.ReplaceAllString(s, "")
	s = strings.Trim(s, "-")
	if s == "" {
		return "unknown"
	}
	return s
}

func Shuffled(rng *rand.Rand) []Fiction {
	n := len(List)
	shuffled := make([]Fiction, n)
	copy(shuffled, List)
	for i := n - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	return shuffled
}
