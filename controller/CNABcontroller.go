package controller

import (
	"desafia-CNAB/services"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CnabController struct {

	TransactionServices services.TransactionServices

}

func (cb *CnabController) UploadFile()gin.HandlerFunc{
	return func(c *gin.Context) {

		file, header, err:= c.Request.FormFile("file")//obtem o arquivo
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//defer file.Close()
		os.Mkdir("./uploads", os.ModePerm)//cria um diretorio, se não existir
		filePath:= "./uploads/" + header.Filename//caminho aonde o arquivo sera salvo
		out, err:= os.Create(filePath)//um novo arquivo para salvar o upload
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro":err.Error()})
			return
		}
		defer out.Close()	
		_, err = io.Copy(out, file)//copia o arquivo vindo do upload para o arquivo criado
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}

		newFile, err:= os.Open(filePath)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"erro":"erro ao reabrir arquivo"})
			return
		}
		defer newFile.Close()

		cb.TransactionServices.ArchiveProcess(newFile)

		c.JSON(http.StatusOK, gin.H{

			"message":"arquivo salvo com sucesso",
			"file": header.Filename,
		})
	}
}

func (cb *CnabController) GetAll()gin.HandlerFunc{
	return func(c *gin.Context) {


		transactions, err:= cb.TransactionServices.TransactionRepository.FindAll()
			if err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
				return
			}

			c.JSON(http.StatusOK, transactions)
		

	}

}



