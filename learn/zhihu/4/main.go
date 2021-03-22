package main

import (
	"./models"
	"github.com/gin-gonic/gin"
	"strconv"
)

//添加操作
func insert(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(200, gin.H{"msg": "name不得为空!"})
		return
	}

	money := c.Query("money")
	if money == "" {
		c.JSON(200, gin.H{"msg": "money不得为空!"})
		return
	}

	Balance, _ := strconv.ParseFloat(money, 64)
	//添加
	user := models.User{}
	user.Name = name
	user.Balance = Balance
	rel, err := models.X.Insert(user)

	if rel == 0 || err != nil {
		c.JSON(200, gin.H{"msg": "添加错误", "err": err, "rel": rel})
	} else {
		c.JSON(200, gin.H{"msg": "添加成功"})
	}
}

//查询单个操作
func get(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(200, gin.H{"msg": "id不得为空!"})
		return
	}
	//string转int64
	ids, _ := strconv.ParseInt(id, 10, 64)
	//查询1
	//user := &User{User_id: ids}
	//rel, err := models.X.Get(user)
	//查询2
	user := &models.User{}
	rel, err := models.X.Where("user_id = ?", ids).Get(user)

	if !rel || err != nil {
		c.JSON(200, gin.H{"msg": "查询错误"})
	} else {
		c.JSON(200, gin.H{"user": user})
	}
}

//查询多条操作
func find(c *gin.Context) {
	users := make(map[int64]models.User)
	err := models.X.Find(&users)
	if err != nil {
		c.JSON(200, gin.H{"msg": err})
	}
	c.JSON(200, gin.H{"msg": users})
}

//修改操作
func updates(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(200, gin.H{"msg": "id1不得为空!", "id": id})
		return
	}
	ids, _ := strconv.ParseInt(id, 10, 64)

	name := c.Query("name")
	if name == "" {
		c.JSON(200, gin.H{"msg": "name不得为空!"})
		return
	}

	balance := c.Query("balance")
	if balance == "" {
		c.JSON(200, gin.H{"msg": "balance不得为空!"})
		return
	}
	money, _ := strconv.ParseFloat(balance, 64)
	//修改
	user := models.User{}
	user.Name = name
	user.Balance = money
	rel, err := models.X.Id(ids).Update(user)

	if rel == 0 || err != nil {
		c.JSON(200, gin.H{"msg": "修改错误!", "rel": rel, "err": err, "user": user})
	} else {
		c.JSON(200, gin.H{"mag": "修改成功"})
	}
}

//删除操作
func delte(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(200, gin.H{"msg": "id1不得为空!", "id": id})
		return
	}
	//string转化int64
	ids, _ := strconv.ParseInt(id, 10, 64)
	//删除
	user := models.User{}
	rel, err := models.X.Id(ids).Delete(user)

	if rel == 0 || err != nil {
		c.JSON(200, gin.H{"msg": "删除错误!", "rel": rel, "err": err, "user": user})
	} else {
		c.JSON(200, gin.H{"mag": "删除成功"})
	}
}

//事务的提交以及回滚
func gun(c *gin.Context) {


	//创建session
	session := models.X.NewSession()
	defer session.Close()
	//创建事务
	err := session.Begin()
	if err != nil {
		c.JSON(200, gin.H{"err": err})
		return
	}
	//操作事务,失败并回滚(模拟购物车结算情景)
	car_id := c.Query("car_id")
	if car_id == "" {
		c.JSON(200, gin.H{"msg": "car_id1不得为空!", "car_id": car_id})
		return
	}

	//查找购物车中的商品id
	ids, _ := strconv.ParseInt(car_id, 10, 64)
	car := &models.Car{Car_id: ids}
	models.X.Get(car)

	/**
	 * goods表库存减去销量
	 */
	//查询商品
	goods := &models.Goods{Goods_id: car.Goods_id}
	models.X.Get(goods)
	//更新库存
	good := models.Goods{}
	good.Stock = goods.Stock - car.Num
	rel4, err4 := session.ID(car.Goods_id).Update(good)
	if rel4 == 0 || err4 != nil {
		session.Rollback()
		c.JSON(200, gin.H{"err4": err4, "rel4": rel4, "carid": car.Goods_id, "goodsid": goods.Goods_id, "Stock": good.Stock})
		return
	}

	/**
	 * 用户扣费
	 */
	//查询用户
	user := &models.User{User_id: car.User_id}
	models.X.Get(user)
	//更新价格
	user_up := models.User{}
	user_up.Balance = user.Balance - car.Total_price
	rel1, err1 := session.ID(car.User_id).Update(user_up)
	if err1 != nil || rel1 == 0 {
		session.Rollback()
		c.JSON(200, gin.H{"err1": err1, "rel1": rel1})
		return
	}

	/**
	 * 删除用户的购物车信息
	 */
	rel2, err2 := session.Delete(car)
	if err2 != nil || rel2 == 0 {
		session.Rollback()
		c.JSON(200, gin.H{"err2": err2, "rel2": rel2})
		return
	}
	if user_up.Balance <= 0 {
		session.Rollback()
		c.JSON(200, gin.H{"msg": "余额不足"})
		return
	}

	err3 := session.Commit()
	if err3 != nil {
		c.JSON(200, gin.H{"err3": err3})
		return
	}
	c.JSON(200, gin.H{"msg": "用户扣费成功"})
}

func update_goods(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(200, gin.H{"msg": "id1不得为空!", "id": id})
		return
	}
	//string转换int64
	ids, err := strconv.ParseInt(id, 10, 64)

	goods_name := c.Query("goods_name")
	if goods_name == "" {
		c.JSON(200, gin.H{"msg": "goods_name不得为空!", "goods_name": goods_name})
		return
	}

	price := c.Query("price")
	if price == "" {
		c.JSON(200, gin.H{"msg": "price不得为空!", "price": price})
		return
	}
	prices, _ := strconv.ParseFloat(price, 64)

	stock := c.Query("stock")
	if stock == "" {
		c.JSON(200, gin.H{"msg": "stock不得为空!", "stock": stock})
		return
	}
	stocks, _ := strconv.ParseInt(stock, 10, 64)

	//修改
	goods := models.Goods{}
	goods.Stock = stocks
	goods.Goods_name = goods_name
	goods.Price = prices
	rel, err := models.X.ID(ids).Update(goods)
	if rel == 0 || err != nil {
		c.JSON(200, gin.H{"msg": "修改失败", "err": err, "stocks": stocks, "goods_name": goods_name, "prices": prices, "id": id})
	} else {
		c.JSON(200, gin.H{"msg": "修改成功"})
	}
}

func shiwu(c *gin.Context) {
	session := models.X.NewSession()
	defer session.Close()

	err := session.Begin()
	user1 := models.User{Name: "xiaoxiao1", Balance: 100}
	_, err = session.Insert(&user1)
	if err != nil {
		return
	}

	session.Rollback()
	data := make(map[string]interface{})
	data["msg"] = "错误"
	c.JSON(200, session)
	c.JSON(200, data)
	return

	//提交
	err = session.Commit()
	if err != nil {
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/insert", insert)
	r.GET("/get", get)
	r.GET("/find", find)
	r.GET("/updates", updates)
	r.GET("/delte", delte)
	r.GET("/update_goods", update_goods)
	r.GET("/gun", gun)
	r.GET("/shiwu", shiwu)
	r.Run(":88")
}