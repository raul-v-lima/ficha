package routes

import (
	"context"
	"fmt"
	"net/http"
	"server/helpers"
	"server/server/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

var batalhaCollection *mongo.Collection = OpenCollection(Client, "batalhas")

func CalcIniciativa() {}
func CalcDano()       {}
func AddPersonagem(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var personagem models.Personagem

	if err := c.BindJSON(&personagem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(personagem)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	personagem.ID = primitive.NewObjectID()
	result, insertErr := batalhaCollection.InsertOne(ctx, personagem)

	if insertErr != nil {
		msg := fmt.Sprintf("personagem não foi criado")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)

}

func CalcularFa(c *gin.Context) int32 {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var magico bool = true

	nome := c.Params.ByName("nome")
	//forca, _ := strconv.ParseInt(c.Params.ByName("forca"), 10, 32)
	forca := helpers.ConvertStringToInt32(c.Params.ByName("forca"))
	destreza := helpers.ConvertStringToInt32(c.Params.ByName("destreza"))
	inteligencia := helpers.ConvertStringToInt32(c.Params.ByName("inteligencia"))

	//intVar, err := strconv.Atoi(strVar)
	var fa int32 = forca + destreza
	if magico {
		fa = inteligencia + destreza
	}

	// limpeza de chamdas de recursos
	defer cancel()
	fmt.Println(fa, nome)
	//c.JSON(http.StatusOK,)
	return fa
}

func UpdatePersonagem(c *gin.Context) {

	personagemID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(personagemID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	type Personagem struct {
		Nome         *string `json:"nome"`
		Forca        *int32  `json:"forca"`
		Armadura     *int32  `json:"armadura"`
		Resistencia  *int32  `json:"resistencia"`
		Destreza     *int32  `json:"destreza"`
		Inteligencia *int32  `json:"inteligencia"`
		Mana         *int32  `json:"mana"`
		Xp           *int32  `json:"xp"`
		Nivel        *int32  `json:"nivel"`
	}
	var personagem Personagem
	if err := c.BindJSON(&personagem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	result, err := batalhaCollection.UpdateOne(ctx, bson.M{"_id": docID}, bson.D{
		{Key: "$set", Value: bson.D{{Key: "nome", Value: personagem.Nome}}},
		{Key: "$set", Value: bson.D{{Key: "forca", Value: personagem.Forca}}},
		{Key: "$set", Value: bson.D{{Key: "armadura", Value: personagem.Armadura}}},
		{Key: "$set", Value: bson.D{{Key: "resistencia", Value: personagem.Resistencia}}},
		{Key: "$set", Value: bson.D{{Key: "destreza", Value: personagem.Destreza}}},
		{Key: "$set", Value: bson.D{{Key: "inteligencia", Value: personagem.Inteligencia}}},
		{Key: "$set", Value: bson.D{{Key: "mana", Value: personagem.Mana}}},
		{Key: "$set", Value: bson.D{{Key: "xp", Value: personagem.Xp}}},
		{Key: "$set", Value: bson.D{{Key: "nivel", Value: personagem.Nivel}}},
	},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.ModifiedCount)

}
func DeletePersonagem(c *gin.Context) {

	personagem := c.Params.ByName("id")
	docId, _ := primitive.ObjectIDFromHex(personagem)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := batalhaCollection.DeleteOne(ctx, bson.M{"_id": docId})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)

}
func Skirmish()          {}
func RelatorioSkirmish() {}
func GetPersonagens(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var personagens []bson.M

	cursor, err := batalhaCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &personagens); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(personagens)

	c.JSON(http.StatusOK, personagens)
}

func GetPersonagem(c *gin.Context) {

	nomePersonagem := c.Params.ByName("nome")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var personagem []bson.M

	cursor, err := batalhaCollection.Find(ctx, bson.M{"nome": nomePersonagem})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &personagem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(personagem)

	c.JSON(http.StatusOK, personagem)
}
