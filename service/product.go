package service

import (
	"context"
	"mginmall/dao"
	"mginmall/model"
	"mginmall/pkg/e"
	"mginmall/pkg/utils"
	"mginmall/serializer"
	"mime/multipart"
	"strconv"
	"sync"

	"gorm.io/gorm"
)

type ProductService struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryID    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	Onsale        bool   `json:"onsale" form:"onsale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

func (service *ProductService) Create(c context.Context, uid uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	if len(files) == 0 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "file 不能为空",
		}
	}
	userDao := dao.NewUserDao(c)
	boss, _ = userDao.GetUserByID(uid)

	tmp, err := files[0].Open()
	if err != nil {
		code = e.ErrorProductImgUpload
		utils.LogrusObj.Errorf("Open Product file:%v", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	path, err := UploadProductToLocalStatic(tmp, uid, service.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		utils.LogrusObj.Errorf("Upload Product To Local Static:%v", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product := model.Product{
		Model:         gorm.Model{},
		Name:          service.Name,
		CategoryID:    service.CategoryID, //CategoryID: ,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           uint(service.Num),
		BossID:        uid,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}

	productDao := dao.NewProductDao(c)
	err = productDao.CreateProduct(&product)

	if err != nil {
		code = e.Error
		utils.LogrusObj.Errorf("Create Product: %v", err)
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductToLocalStatic(tmp, uid, service.Name+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(&product),
	}
}

func (service *ProductService) List(c context.Context) serializer.Response {
	var products []model.Product
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	condition := make(map[string]interface{})
	if service.CategoryID != 0 {
		condition["category_id"] = service.CategoryID
	}

	productDao := dao.NewProductDao(c)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, &service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

func (service *ProductService) Search(c context.Context) serializer.Response {
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewProductDao(c)
	products, count, err := productDao.SearchProDuct(service.Info, service.BasePage) //ES实现search
	if err != nil {
		code = e.Error
		utils.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(count))
}

func (service *ProductService) Show(c context.Context, id string) serializer.Response {
	code := e.Success
	pid, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(c)
	product, err := productDao.GetProductByID(uint(pid))
	if err != nil {
		code = e.Error
		utils.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}

}
