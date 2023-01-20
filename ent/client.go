// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"shopular/ent/migrate"

	"shopular/ent/cart"
	"shopular/ent/category"
	"shopular/ent/customer"
	"shopular/ent/order"
	"shopular/ent/product"
	"shopular/ent/subcategory"
	"shopular/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Cart is the client for interacting with the Cart builders.
	Cart *CartClient
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
	// Product is the client for interacting with the Product builders.
	Product *ProductClient
	// SubCategory is the client for interacting with the SubCategory builders.
	SubCategory *SubCategoryClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Cart = NewCartClient(c.config)
	c.Category = NewCategoryClient(c.config)
	c.Customer = NewCustomerClient(c.config)
	c.Order = NewOrderClient(c.config)
	c.Product = NewProductClient(c.config)
	c.SubCategory = NewSubCategoryClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Cart:        NewCartClient(cfg),
		Category:    NewCategoryClient(cfg),
		Customer:    NewCustomerClient(cfg),
		Order:       NewOrderClient(cfg),
		Product:     NewProductClient(cfg),
		SubCategory: NewSubCategoryClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Cart:        NewCartClient(cfg),
		Category:    NewCategoryClient(cfg),
		Customer:    NewCustomerClient(cfg),
		Order:       NewOrderClient(cfg),
		Product:     NewProductClient(cfg),
		SubCategory: NewSubCategoryClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Cart.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Cart.Use(hooks...)
	c.Category.Use(hooks...)
	c.Customer.Use(hooks...)
	c.Order.Use(hooks...)
	c.Product.Use(hooks...)
	c.SubCategory.Use(hooks...)
	c.User.Use(hooks...)
}

// CartClient is a client for the Cart schema.
type CartClient struct {
	config
}

// NewCartClient returns a client for the Cart from the given config.
func NewCartClient(c config) *CartClient {
	return &CartClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cart.Hooks(f(g(h())))`.
func (c *CartClient) Use(hooks ...Hook) {
	c.hooks.Cart = append(c.hooks.Cart, hooks...)
}

// Create returns a builder for creating a Cart entity.
func (c *CartClient) Create() *CartCreate {
	mutation := newCartMutation(c.config, OpCreate)
	return &CartCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cart entities.
func (c *CartClient) CreateBulk(builders ...*CartCreate) *CartCreateBulk {
	return &CartCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cart.
func (c *CartClient) Update() *CartUpdate {
	mutation := newCartMutation(c.config, OpUpdate)
	return &CartUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CartClient) UpdateOne(ca *Cart) *CartUpdateOne {
	mutation := newCartMutation(c.config, OpUpdateOne, withCart(ca))
	return &CartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CartClient) UpdateOneID(id int) *CartUpdateOne {
	mutation := newCartMutation(c.config, OpUpdateOne, withCartID(id))
	return &CartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cart.
func (c *CartClient) Delete() *CartDelete {
	mutation := newCartMutation(c.config, OpDelete)
	return &CartDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CartClient) DeleteOne(ca *Cart) *CartDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CartClient) DeleteOneID(id int) *CartDeleteOne {
	builder := c.Delete().Where(cart.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CartDeleteOne{builder}
}

// Query returns a query builder for Cart.
func (c *CartClient) Query() *CartQuery {
	return &CartQuery{
		config: c.config,
	}
}

// Get returns a Cart entity by its id.
func (c *CartClient) Get(ctx context.Context, id int) (*Cart, error) {
	return c.Query().Where(cart.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CartClient) GetX(ctx context.Context, id int) *Cart {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Cart.
func (c *CartClient) QueryUser(ca *Cart) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cart.Table, cart.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, cart.UserTable, cart.UserPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProduct queries the product edge of a Cart.
func (c *CartClient) QueryProduct(ca *Cart) *ProductQuery {
	query := &ProductQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cart.Table, cart.FieldID, id),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, cart.ProductTable, cart.ProductPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CartClient) Hooks() []Hook {
	return c.hooks.Cart
}

// CategoryClient is a client for the Category schema.
type CategoryClient struct {
	config
}

// NewCategoryClient returns a client for the Category from the given config.
func NewCategoryClient(c config) *CategoryClient {
	return &CategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `category.Hooks(f(g(h())))`.
func (c *CategoryClient) Use(hooks ...Hook) {
	c.hooks.Category = append(c.hooks.Category, hooks...)
}

// Create returns a builder for creating a Category entity.
func (c *CategoryClient) Create() *CategoryCreate {
	mutation := newCategoryMutation(c.config, OpCreate)
	return &CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Category entities.
func (c *CategoryClient) CreateBulk(builders ...*CategoryCreate) *CategoryCreateBulk {
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Category.
func (c *CategoryClient) Update() *CategoryUpdate {
	mutation := newCategoryMutation(c.config, OpUpdate)
	return &CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CategoryClient) UpdateOne(ca *Category) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategory(ca))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CategoryClient) UpdateOneID(id int) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategoryID(id))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Category.
func (c *CategoryClient) Delete() *CategoryDelete {
	mutation := newCategoryMutation(c.config, OpDelete)
	return &CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CategoryClient) DeleteOne(ca *Category) *CategoryDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CategoryClient) DeleteOneID(id int) *CategoryDeleteOne {
	builder := c.Delete().Where(category.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CategoryDeleteOne{builder}
}

// Query returns a query builder for Category.
func (c *CategoryClient) Query() *CategoryQuery {
	return &CategoryQuery{
		config: c.config,
	}
}

// Get returns a Category entity by its id.
func (c *CategoryClient) Get(ctx context.Context, id int) (*Category, error) {
	return c.Query().Where(category.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CategoryClient) GetX(ctx context.Context, id int) *Category {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySubs queries the subs edge of a Category.
func (c *CategoryClient) QuerySubs(ca *Category) *SubCategoryQuery {
	query := &SubCategoryQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(subcategory.Table, subcategory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, category.SubsTable, category.SubsColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CategoryClient) Hooks() []Hook {
	return c.hooks.Category
}

// CustomerClient is a client for the Customer schema.
type CustomerClient struct {
	config
}

// NewCustomerClient returns a client for the Customer from the given config.
func NewCustomerClient(c config) *CustomerClient {
	return &CustomerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customer.Hooks(f(g(h())))`.
func (c *CustomerClient) Use(hooks ...Hook) {
	c.hooks.Customer = append(c.hooks.Customer, hooks...)
}

// Create returns a builder for creating a Customer entity.
func (c *CustomerClient) Create() *CustomerCreate {
	mutation := newCustomerMutation(c.config, OpCreate)
	return &CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Customer entities.
func (c *CustomerClient) CreateBulk(builders ...*CustomerCreate) *CustomerCreateBulk {
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Customer.
func (c *CustomerClient) Update() *CustomerUpdate {
	mutation := newCustomerMutation(c.config, OpUpdate)
	return &CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerClient) UpdateOne(cu *Customer) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomer(cu))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerClient) UpdateOneID(id int) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomerID(id))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Customer.
func (c *CustomerClient) Delete() *CustomerDelete {
	mutation := newCustomerMutation(c.config, OpDelete)
	return &CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CustomerClient) DeleteOne(cu *Customer) *CustomerDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CustomerClient) DeleteOneID(id int) *CustomerDeleteOne {
	builder := c.Delete().Where(customer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerDeleteOne{builder}
}

// Query returns a query builder for Customer.
func (c *CustomerClient) Query() *CustomerQuery {
	return &CustomerQuery{
		config: c.config,
	}
}

// Get returns a Customer entity by its id.
func (c *CustomerClient) Get(ctx context.Context, id int) (*Customer, error) {
	return c.Query().Where(customer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerClient) GetX(ctx context.Context, id int) *Customer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Customer.
func (c *CustomerClient) QueryUser(cu *Customer) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, customer.UserTable, customer.UserColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CustomerClient) Hooks() []Hook {
	return c.hooks.Customer
}

// OrderClient is a client for the Order schema.
type OrderClient struct {
	config
}

// NewOrderClient returns a client for the Order from the given config.
func NewOrderClient(c config) *OrderClient {
	return &OrderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `order.Hooks(f(g(h())))`.
func (c *OrderClient) Use(hooks ...Hook) {
	c.hooks.Order = append(c.hooks.Order, hooks...)
}

// Create returns a builder for creating a Order entity.
func (c *OrderClient) Create() *OrderCreate {
	mutation := newOrderMutation(c.config, OpCreate)
	return &OrderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Order entities.
func (c *OrderClient) CreateBulk(builders ...*OrderCreate) *OrderCreateBulk {
	return &OrderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Order.
func (c *OrderClient) Update() *OrderUpdate {
	mutation := newOrderMutation(c.config, OpUpdate)
	return &OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderClient) UpdateOne(o *Order) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrder(o))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderClient) UpdateOneID(id int) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrderID(id))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Order.
func (c *OrderClient) Delete() *OrderDelete {
	mutation := newOrderMutation(c.config, OpDelete)
	return &OrderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrderClient) DeleteOne(o *Order) *OrderDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrderClient) DeleteOneID(id int) *OrderDeleteOne {
	builder := c.Delete().Where(order.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderDeleteOne{builder}
}

// Query returns a query builder for Order.
func (c *OrderClient) Query() *OrderQuery {
	return &OrderQuery{
		config: c.config,
	}
}

// Get returns a Order entity by its id.
func (c *OrderClient) Get(ctx context.Context, id int) (*Order, error) {
	return c.Query().Where(order.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderClient) GetX(ctx context.Context, id int) *Order {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OrderClient) Hooks() []Hook {
	return c.hooks.Order
}

// ProductClient is a client for the Product schema.
type ProductClient struct {
	config
}

// NewProductClient returns a client for the Product from the given config.
func NewProductClient(c config) *ProductClient {
	return &ProductClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `product.Hooks(f(g(h())))`.
func (c *ProductClient) Use(hooks ...Hook) {
	c.hooks.Product = append(c.hooks.Product, hooks...)
}

// Create returns a builder for creating a Product entity.
func (c *ProductClient) Create() *ProductCreate {
	mutation := newProductMutation(c.config, OpCreate)
	return &ProductCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Product entities.
func (c *ProductClient) CreateBulk(builders ...*ProductCreate) *ProductCreateBulk {
	return &ProductCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Product.
func (c *ProductClient) Update() *ProductUpdate {
	mutation := newProductMutation(c.config, OpUpdate)
	return &ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProductClient) UpdateOne(pr *Product) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProduct(pr))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProductClient) UpdateOneID(id int) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProductID(id))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Product.
func (c *ProductClient) Delete() *ProductDelete {
	mutation := newProductMutation(c.config, OpDelete)
	return &ProductDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProductClient) DeleteOne(pr *Product) *ProductDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProductClient) DeleteOneID(id int) *ProductDeleteOne {
	builder := c.Delete().Where(product.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProductDeleteOne{builder}
}

// Query returns a query builder for Product.
func (c *ProductClient) Query() *ProductQuery {
	return &ProductQuery{
		config: c.config,
	}
}

// Get returns a Product entity by its id.
func (c *ProductClient) Get(ctx context.Context, id int) (*Product, error) {
	return c.Query().Where(product.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProductClient) GetX(ctx context.Context, id int) *Product {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySub queries the sub edge of a Product.
func (c *ProductClient) QuerySub(pr *Product) *SubCategoryQuery {
	query := &SubCategoryQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, id),
			sqlgraph.To(subcategory.Table, subcategory.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, product.SubTable, product.SubPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCart queries the cart edge of a Product.
func (c *ProductClient) QueryCart(pr *Product) *CartQuery {
	query := &CartQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, id),
			sqlgraph.To(cart.Table, cart.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, product.CartTable, product.CartPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProductClient) Hooks() []Hook {
	return c.hooks.Product
}

// SubCategoryClient is a client for the SubCategory schema.
type SubCategoryClient struct {
	config
}

// NewSubCategoryClient returns a client for the SubCategory from the given config.
func NewSubCategoryClient(c config) *SubCategoryClient {
	return &SubCategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `subcategory.Hooks(f(g(h())))`.
func (c *SubCategoryClient) Use(hooks ...Hook) {
	c.hooks.SubCategory = append(c.hooks.SubCategory, hooks...)
}

// Create returns a builder for creating a SubCategory entity.
func (c *SubCategoryClient) Create() *SubCategoryCreate {
	mutation := newSubCategoryMutation(c.config, OpCreate)
	return &SubCategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SubCategory entities.
func (c *SubCategoryClient) CreateBulk(builders ...*SubCategoryCreate) *SubCategoryCreateBulk {
	return &SubCategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SubCategory.
func (c *SubCategoryClient) Update() *SubCategoryUpdate {
	mutation := newSubCategoryMutation(c.config, OpUpdate)
	return &SubCategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubCategoryClient) UpdateOne(sc *SubCategory) *SubCategoryUpdateOne {
	mutation := newSubCategoryMutation(c.config, OpUpdateOne, withSubCategory(sc))
	return &SubCategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubCategoryClient) UpdateOneID(id int) *SubCategoryUpdateOne {
	mutation := newSubCategoryMutation(c.config, OpUpdateOne, withSubCategoryID(id))
	return &SubCategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SubCategory.
func (c *SubCategoryClient) Delete() *SubCategoryDelete {
	mutation := newSubCategoryMutation(c.config, OpDelete)
	return &SubCategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SubCategoryClient) DeleteOne(sc *SubCategory) *SubCategoryDeleteOne {
	return c.DeleteOneID(sc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SubCategoryClient) DeleteOneID(id int) *SubCategoryDeleteOne {
	builder := c.Delete().Where(subcategory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubCategoryDeleteOne{builder}
}

// Query returns a query builder for SubCategory.
func (c *SubCategoryClient) Query() *SubCategoryQuery {
	return &SubCategoryQuery{
		config: c.config,
	}
}

// Get returns a SubCategory entity by its id.
func (c *SubCategoryClient) Get(ctx context.Context, id int) (*SubCategory, error) {
	return c.Query().Where(subcategory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubCategoryClient) GetX(ctx context.Context, id int) *SubCategory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCategory queries the category edge of a SubCategory.
func (c *SubCategoryClient) QueryCategory(sc *SubCategory) *CategoryQuery {
	query := &CategoryQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := sc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subcategory.Table, subcategory.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subcategory.CategoryTable, subcategory.CategoryColumn),
		)
		fromV = sqlgraph.Neighbors(sc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProduct queries the product edge of a SubCategory.
func (c *SubCategoryClient) QueryProduct(sc *SubCategory) *ProductQuery {
	query := &ProductQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := sc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subcategory.Table, subcategory.FieldID, id),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, subcategory.ProductTable, subcategory.ProductPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(sc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubCategoryClient) Hooks() []Hook {
	return c.hooks.SubCategory
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCarts queries the carts edge of a User.
func (c *UserClient) QueryCarts(u *User) *CartQuery {
	query := &CartQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(cart.Table, cart.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.CartsTable, user.CartsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCustomer queries the customer edge of a User.
func (c *UserClient) QueryCustomer(u *User) *CustomerQuery {
	query := &CustomerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.CustomerTable, user.CustomerColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
