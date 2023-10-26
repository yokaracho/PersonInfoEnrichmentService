package handler

import (
	"NameService/pkg/api"
	"NameService/pkg/logger"
	"NameService/pkg/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var getLogger = logger.GetLogger()

func (h *Handler) InsertData(c echo.Context) error {
	fio := &model.NameModel{}

	if err := json.NewDecoder(c.Request().Body).Decode(fio); err != nil {
		getLogger.Infof("Неверные данные в запросе: %s", err.Error())
		getLogger.Debugf("Неверные данные в запросе: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверные данные в запросе"})
	}
	age, gender, nationality, err := api.EnrichData(fio)
	if err != nil {
		getLogger.Infof("Ошибка при обогащении данных: %s", err.Error())
		getLogger.Debugf("Ошибка при обогащении данных: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при обогащении данных"})
	}

	fio.Age = age
	fio.Gender = gender
	fio.Nationality = nationality
	rowsAffected, err := h.service.InsertData(c.Request().Context(), fio)
	if err != nil {
		fmt.Println("Ошибка при вставке данных:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при вставке данных"})
	}

	getLogger.Infof("Данные вставлены, количество строк: %d", rowsAffected)
	getLogger.Debugf("Данные вставлены, количество строк: %d", rowsAffected)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Данные вставлены, количество строк: %d", rowsAffected),
	})
}

func (h *Handler) GetPeopleByAge(c echo.Context) error {
	people, err := h.service.GetPeopleByAge(context.Background())
	if err != nil {
		getLogger.Infof("Ошибка: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error: " + err.Error()})
	}
	getLogger.Debugf("Выводим все данные")
	getLogger.Infof("Выводим все данные")
	return c.JSON(http.StatusOK, people)
}

//	func (h *Handler) GetAllImages(c echo.Context) error {
//		images, err := h.service.GetAllImages(context.Background())
//		if err != nil {
//			log.Printf("Ошибка:&v", err)
//			return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
//		}
//
//		return c.JSON(http.StatusOK, images)
//	}

func (h *Handler) UpdateInfo(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		getLogger.Infof("Неверный идентификатор: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный идентификатор"})
	}

	// Извлечение новых данных о человеке из запроса
	var updateData model.NameModel
	if err := c.Bind(&updateData); err != nil {
		getLogger.Debugf("Ошибка при извлечении данных из запроса: %s", err.Error())
		getLogger.Infof("Ошибка при извлечении данных из запроса: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверные данные в запросе"})
	}
	updateData.ID = id

	getLogger.Debugf("Обновляем информацию для ID: %d, новые данные: %+v", updateData.ID, updateData)

	if err := h.service.UpdateInfo(ctx, &updateData); err != nil {
		getLogger.Debugf("Ошибка при обновлении информации: %s", err.Error())
		getLogger.Infof("Ошибка при обновлении информации: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при обновлении информации"})
	}

	getLogger.Debugf("Информация успешно обновлена для ID: %d", updateData.ID)
	getLogger.Infof("Информация успешно обновлена для ID: %d", updateData.ID)
	return c.JSON(http.StatusOK, map[string]string{"message": "Информация обновлена"})
}

func (h *Handler) DeleteImageByID(c echo.Context) error {
	// Получите ID из параметров маршрута
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		getLogger.Debugf("Не совпадает ID:%s", err.Error())
		getLogger.Infof("Не совпадает ID:%s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	// Вызовите метод сервиса для удаления записи по ID
	rowsAffected, err := h.service.DeleteImageByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if rowsAffected == 0 {
		getLogger.Infof("Запись не найдена")
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Record not found"})
	}

	getLogger.Debugf("Запись удалена")
	getLogger.Infof("Запись удалена")
	return c.JSON(http.StatusOK, map[string]string{"message": "Record deleted"})
}
