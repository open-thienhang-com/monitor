var spec = {
  openapi: "3.0.1",
  info: {
    version: "1.0.0",
    title: "Swagger Fresh Retail",
  },
  paths: {
    "/v1/inventory/product/{id}": {
      get: {
        summary: "Find product by ID",
        description: "Return information of product",
        operationId: "getProductById",
        tags: ["Product"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get product success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ProductResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update product information.",
        operationId: "updateProduct",
        tags: ["Product"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update product information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/ProductNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/ProductNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/ProductNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update product information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ProductResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete product by ID",
        description: "Delete a product out of systems",
        operationId: "deleteProductById",
        tags: ["Product"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete product success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/product": {
      post: {
        summary: "Add a new product to the systems",
        description: "Add a new product to the systems",
        operationId: "addProduct",
        tags: ["Product"],
        requestBody: {
          description: "Create a new product to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/ProductNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/ProductNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/ProductNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new product to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ProductResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/products": {
      get: {
        summary: "Get list product",
        description: "Return list product",
        operationId: "getListProduct",
        tags: ["Product"],
        // parameters: [
        //   {
        //     $ref: "#/components/parameters/Id",
        //   },
        // ],
        responses: {
          200: {
            description: "Get list product success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListProductResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/product/search": {
      get: {
        summary: "Search product by name",
        description: "Return list product",
        operationId: "getListProduct",
        tags: ["Product"],
        parameters: [{
            $ref: "#/components/parameters/Name",
          },
          {
            $ref: "#/components/parameters/Description",
          },
        ],
        responses: {
          200: {
            description: "Search product success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListProductResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/store/{id}": {
      get: {
        summary: "Find store by ID",
        description: "Return information of store",
        operationId: "getStoreById",
        tags: ["Store"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get store success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/StoreResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update store information.",
        operationId: "updateStore",
        tags: ["Store"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update store information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/StoreNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/StoreNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/StoreNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update store information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/StoreResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete store by ID",
        description: "Delete a store out of systems",
        operationId: "deleteProductById",
        tags: ["Store"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete store success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/store": {
      post: {
        summary: "Add a new store to the systems",
        description: "Add a new store to the systems",
        operationId: "addStore",
        tags: ["Store"],
        requestBody: {
          description: "Create a new store to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/StoreNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/StoreNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/StoreNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new store to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/StoreResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/store/findbyname": {
      get: {
        summary: "Search store by name",
        description: "Return list store",
        operationId: "searchStore",
        tags: ["Store"],
        parameters: [{
          $ref: "#/components/parameters/Name",
        }, ],
        responses: {
          200: {
            description: "Search store success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListStoreResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/price/{id}": {
      get: {
        summary: "Find price by ID",
        description: "Return information of price product",
        operationId: "getPriceById",
        tags: ["Price"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get price success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/PriceResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update price information.",
        operationId: "updatePrice",
        tags: ["Price"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update price information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/PriceNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/PriceNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/PriceNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update price information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/PriceResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete price by ID",
        description: "Delete price of product out of systems",
        operationId: "deletePriceById",
        tags: ["Price"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete price success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/price": {
      post: {
        summary: "Add a new price to the systems",
        description: "Add a new price to the systems",
        operationId: "addPrice",
        tags: ["Price"],
        requestBody: {
          description: "Create a new price to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/PriceNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/PriceNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/PriceNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new price to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/PriceResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/price/findbyproductid": {
      get: {
        summary: "Search price by id product",
        description: "Return list price",
        operationId: "searchPrice",
        tags: ["Price"],
        parameters: [{
          $ref: "#/components/parameters/ProductID",
        }, ],
        responses: {
          200: {
            description: "Search price success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListPriceResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/campaign/{id}": {
      get: {
        summary: "Find Campaign by ID",
        description: "Return information of price campaign",
        operationId: "getCampaignById",
        tags: ["Campaign"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get campaign success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/CampaignResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update campaign information.",
        operationId: "updateCampaign",
        tags: ["Campaign"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update campaign information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/CampaignNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/CampaignNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/CampaignNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update campaign information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/CampaignResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete campaign by ID",
        description: "Delete campaign out of systems",
        operationId: "deleteCampaignById",
        tags: ["Campaign"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete campaign success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/campaign": {
      post: {
        summary: "Add a new campaign to the systems",
        description: "Add a new campaign to the systems",
        operationId: "addCampaign",
        tags: ["Campaign"],
        requestBody: {
          description: "Create a new campaign to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/CampaignNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/CampaignNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/CampaignNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new campaign to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/CampaignResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/campaign/findbynamecampaign": {
      get: {
        summary: "Search campaign by name",
        description: "Return list campaign",
        operationId: "searchCampaign",
        tags: ["Campaign"],
        parameters: [{
          $ref: "#/components/parameters/Name",
        }, ],
        responses: {
          200: {
            description: "Search campaign success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListCampaignResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/customer/{id}": {
      get: {
        summary: "Find Customer by ID",
        description: "Return information of customer",
        operationId: "getCustomerById",
        tags: ["Customer"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get price success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/CustomerResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update Customer information.",
        operationId: "updateCustomer",
        tags: ["Customer"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update customer information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/CustomerNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/CustomerNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/CustomerNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update price information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/PriceResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete price by ID",
        description: "Delete price of product out of systems",
        operationId: "deletePriceById",
        tags: ["Customer"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete price success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/customer": {
      post: {
        summary: "Add a new customer to the systems",
        description: "Add a new customer to the systems",
        operationId: "addCustomer",
        tags: ["Customer"],
        requestBody: {
          description: "Create a new customer to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/CustomerNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/CustomerNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/CustomerNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new customer to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/CustomerResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/customer/findbynamecustomer": {
      get: {
        summary: "Search customer by id product",
        description: "Return list customer",
        operationId: "searchCustomer",
        tags: ["Customer"],
        parameters: [{
          $ref: "#/components/parameters/Name",
        }, ],
        responses: {
          200: {
            description: "Search customer success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListCustomerResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/date_order/{id}": {
      get: {
        summary: "Find date order by ID",
        description: "Return information of DateOrder",
        operationId: "getDateOrderById",
        tags: ["DateOrder"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get date order success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DateOrderResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update date order information.",
        operationId: "updateDateOrder",
        tags: ["DateOrder"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update date order information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/DateOrderNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/DateOrderNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/DateOrderNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update date order information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DateOrderResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete date order by ID",
        description: "Delete date order out of systems",
        operationId: "deleteDateOrderById",
        tags: ["DateOrder"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete date order success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/date_order": {
      post: {
        summary: "Add a new date order to the systems",
        description: "Add a new date order to the systems",
        operationId: "addPrice",
        tags: ["DateOrder"],
        requestBody: {
          description: "Create a new date order to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/DateOrderNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/DateOrderNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/DateOrderNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new date order to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DateOrderResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/delivery/{id}": {
      get: {
        summary: "Find Delivery by ID",
        description: "Return information of delivery",
        operationId: "getDeliveryById",
        tags: ["Delivery"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get delivery success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeliveryResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update delivery information.",
        operationId: "updateDelivery",
        tags: ["Delivery"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update delivery information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/DeliveryNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/DeliveryNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/DeliveryNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update delivery information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeliveryResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete delivery by ID",
        description: "Delete delivery out of systems",
        operationId: "deleteDeliveryById",
        tags: ["Delivery"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete delivery success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/delivery": {
      post: {
        summary: "Add a new delivery to the systems",
        description: "Add a new delivery to the systems",
        operationId: "addDelivery",
        tags: ["Delivery"],
        requestBody: {
          description: "Create a new delivery to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/DeliveryNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/DeliveryNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/DeliveryNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new Delivery to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeliveryResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/delivery/findbymethod": {
      get: {
        summary: "Search delivery by method",
        description: "Return list delivery",
        operationId: "searchDelivery",
        tags: ["Delivery"],
        parameters: [{
          $ref: "#/components/parameters/Method",
        }, ],
        responses: {
          200: {
            description: "Search delivery success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListDeliveryResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/discount/{id}": {
      get: {
        summary: "Find discount by ID",
        description: "Return information of discount",
        operationId: "getDiscountById",
        tags: ["Discount"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get discount success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DiscountResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update discount information.",
        operationId: "updateDiscount",
        tags: ["Discount"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update discount information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/DiscountNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/DiscountNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/DiscountNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update discount information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DiscountResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete discount by ID",
        description: "Delete discount out of systems",
        operationId: "deleteDiscountById",
        tags: ["Discount"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete discount success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/discount": {
      post: {
        summary: "Add a new discount to the systems",
        description: "Add a new discount to the systems",
        operationId: "addDiscount",
        tags: ["Discount"],
        requestBody: {
          description: "Create a new discount to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/DiscountNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/DiscountNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/DiscountNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new discount to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DiscountResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/discounts": {
      get: {
        summary: "Get list discount",
        description: "Return list discount",
        operationId: "getListDiscount",
        tags: ["Discount"],
        responses: {
          200: {
            description: "Get list discount success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListDiscountResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/discount_type/{id}": {
      get: {
        summary: "Find Discount Type by ID",
        description: "Return information of Discount Type",
        operationId: "getPriceById",
        tags: ["DiscountType"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get Discount Type success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DiscountTypeResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update Discount Type information.",
        operationId: "updateDiscountType",
        tags: ["DiscountType"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update Discount Type information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/DiscountTypeNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/DiscountTypeNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/DiscountTypeNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update Discount Type information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DiscountTypeResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete Discount Type by ID",
        description: "Delete Discount Type out of systems",
        operationId: "deletePriceById",
        tags: ["DiscountType"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete Discount Type success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/discount_type": {
      post: {
        summary: "Add a new Discount Type to the systems",
        description: "Add a new Discount Type to the systems",
        operationId: "addPrice",
        tags: ["DiscountType"],
        requestBody: {
          description: "Create a new Discount Type to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/DiscountTypeNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/DiscountTypeNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/DiscountTypeNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new Discount Type to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DiscountTypeResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/discount_types": {
      get: {
        summary: "Get list Discount Type",
        description: "Return list DiscountType",
        operationId: "getListDiscountType",
        tags: ["DiscountType"],
        responses: {
          200: {
            description: "Get list Discount Type success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListDiscountTypeResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/discount_value/{id}": {
      get: {
        summary: "Find Discount Value by ID",
        description: "Return information of Discount Value",
        operationId: "getDiscountValueById",
        tags: ["DiscountValue"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get Discount Value success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DiscountValueResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update Discount Value information.",
        operationId: "updateDiscountValue",
        tags: ["DiscountValue"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update Discount Value information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/DiscountValueNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/DiscountValueNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/DiscountValueNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update Discount Value information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DiscountValueResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete DiscountValue by ID",
        description: "Delete DiscountValue out of systems",
        operationId: "deleteDiscountValueById",
        tags: ["DiscountValue"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete Discount Value success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/discount_value": {
      post: {
        summary: "Add a new Discount Value to the systems",
        description: "Add a new Discount Value to the systems",
        operationId: "addDiscountValue",
        tags: ["DiscountValue"],
        requestBody: {
          description: "Create a new Discount Value to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/DiscountValueNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/DiscountValueNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/DiscountValueNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new Discount Value to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DiscountValueResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/discount_values": {
      get: {
        summary: "Get list Discount Value",
        description: "Return list Discount Value",
        operationId: "getListDiscountValue",
        tags: ["DiscountValue"],
        responses: {
          200: {
            description: "Get list Discount Value success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListDiscountValueResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/discount_value/findbyname": {
      get: {
        summary: "Search Discount Value by display name",
        description: "Return list Discount Value",
        operationId: "searchDiscountValue",
        tags: ["DiscountValue"],
        parameters: [{
          $ref: "#/components/parameters/Name",
        }, ],
        responses: {
          200: {
            description: "Search Discount Value success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListDiscountValueResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/employee/{id}": {
      get: {
        summary: "Find employee by ID",
        description: "Return information of employee",
        operationId: "getEmployeeById",
        tags: ["Employee"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get employee success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/EmployeeResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update employee information.",
        operationId: "updateEmployee",
        tags: ["Employee"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update employee information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/EmployeeNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/EmployeeNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/EmployeeNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update employee information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/EmployeeResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete employee by ID",
        description: "Delete employee out of systems",
        operationId: "deletePriceById",
        tags: ["Employee"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete employee success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/employee": {
      post: {
        summary: "Add a new employee to the systems",
        description: "Add a new employee to the systems",
        operationId: "addPrice",
        tags: ["Employee"],
        requestBody: {
          description: "Create a new employee to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/EmployeeNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/EmployeeNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/EmployeeNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new employee to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/PriceResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/employee/findbyname": {
      get: {
        summary: "Search employee by name employee",
        description: "Return list employee",
        operationId: "searchPrice",
        tags: ["Employee"],
        parameters: [{
          $ref: "#/components/parameters/Name",
        }, ],
        responses: {
          200: {
            description: "Search employee success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListEmployeeResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/order/{id}": {
      get: {
        summary: "Find order by ID",
        description: "Return information of order",
        operationId: "getOrderById",
        tags: ["Order"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get order success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/OrderResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update order information.",
        operationId: "updateOrder",
        tags: ["Order"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update order information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/OrderNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/OrderNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/OrderNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update order information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/OrderResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete order by ID",
        description: "Delete order out of systems",
        operationId: "deleteOrderById",
        tags: ["Order"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete order success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/order": {
      post: {
        summary: "Add a new order to the systems",
        description: "Add a new order to the systems",
        operationId: "addOrder",
        tags: ["Order"],
        requestBody: {
          description: "Create a new order to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/OrderNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/OrderNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/OrderNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new order to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/OrderResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/orders": {
      get: {
        summary: "Get list order",
        description: "Return list order",
        operationId: "getListOrder",
        tags: ["Order"],
        responses: {
          200: {
            description: "Get list order success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListOrderResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/payment/{id}": {
      get: {
        summary: "Find payment by ID",
        description: "Return information payment",
        operationId: "getPaymentById",
        tags: ["Payment"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get payment success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/PaymentResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update payment information.",
        operationId: "updatePayment",
        tags: ["Payment"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update payment information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/PaymentNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/PaymentNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/PaymentNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update payment information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/PaymentResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete payment by ID",
        description: "Delete payment out of systems",
        operationId: "deletePaymentById",
        tags: ["Payment"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete payment success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/payment": {
      post: {
        summary: "Add a new payment to the systems",
        description: "Add a new payment to the systems",
        operationId: "addPayment",
        tags: ["Payment"],
        requestBody: {
          description: "Create a new payment to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/PaymentNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/PaymentNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/PaymentNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new payment to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/PaymentResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/payment/findbymethod": {
      get: {
        summary: "Search payment by method",
        description: "Return list payment",
        operationId: "searchPayment",
        tags: ["Payment"],
        parameters: [{
          $ref: "#/components/parameters/Method",
        }, ],
        responses: {
          200: {
            description: "Search payment success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListPaymentResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/strategy/{id}": {
      get: {
        summary: "Find Strategy by ID",
        description: "Return information of strategy",
        operationId: "getStrategyById",
        tags: ["Strategy"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Get strategy success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/StrategyResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      put: {
        summary: "Update strategy information.",
        operationId: "updateStrategy",
        tags: ["Strategy"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        requestBody: {
          description: "Update strategy information in the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/StrategyNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/StrategyNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/StrategyNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Update strategy information in the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/StrategyResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
      delete: {
        summary: "Delete strategy by ID",
        description: "Delete strategy out of systems",
        operationId: "deleteStrategyById",
        tags: ["Strategy"],
        parameters: [{
          $ref: "#/components/parameters/Id",
        }, ],
        responses: {
          200: {
            description: "Delete strategy success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/DeleteSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/strategy": {
      post: {
        summary: "Add a new strategy to the systems",
        description: "Add a new strategy to the systems",
        operationId: "addPrice",
        tags: ["Strategy"],
        requestBody: {
          description: "Create a new strategy to the systems",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/parameters/StrategyNoID",
              },
            },
            "application/xml": {
              schema: {
                $ref: "#/components/parameters/StrategyNoID",
              },
            },
            "application/x-www-form-urlencoded": {
              schema: {
                $ref: "#/components/parameters/StrategyNoID",
              },
            },
          },
          required: true,
        },
        responses: {
          200: {
            description: "Create a new strategy to the systems, Success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/StrategyResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
    "/v1/inventory/strategy/findbynamestrategy": {
      get: {
        summary: "Search strategy by name strategy",
        description: "Return list strategy",
        operationId: "searchStrategy",
        tags: ["Strategy"],
        parameters: [{
          $ref: "#/components/parameters/Name",
        }, ],
        responses: {
          200: {
            description: "Search strategy success",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/parameters/Success/ListStrategyResponeSuccess",
                },
              },
            },
          },
          404: {
            $ref: "#/components/responses/Error404",
          },
          400: {
            $ref: "#/components/responses/Error400",
          },
        },
      },
    },
  },
  components: {
    schemas: {
      Product: {
        required: [
          "name",
          "description",
          "price",
          "stock",
          "unit",
          "category",
          "brand",
          "sale_status",
          "sales_type",
        ],
        properties: {
          id: {
            description: "id of the product",
            type: "number",
            example: 5,
          },
          name: {
            description: "Name of the product",
            type: "string",
            example: "Apple",
          },
          description: {
            description: "Description of the product",
            type: "string",
            example: "This is fresh food.",
          },
          price: {
            description: "Price of the product",
            type: "number",
            format: "float",
            example: 19.99,
          },
          stock: {
            description: "Stock quantity of the product",
            type: "integer",
            example: 100,
          },
          unit: {
            description: "Unit of product volume",
            type: "string",
            example: "kg",
          },
          category: {
            description: "Category of the product",
            type: "string",
            example: "Fruit",
          },
          brand: {
            description: "Brand of the product",
            type: "string",
            example: "NewZealand",
          },
          sale_status: {
            description: "Product sales status",
            type: "string",
            example: "In Stock",
          },
          sales_type: {
            description: "Sales type of the product",
            type: "string",
            example: "Online",
          },
        },
      },
      Store: {
        required: [
          "name",
          "city",
          "country",
          "street_name",
          "street_number",
          "manager",
        ],
        properties: {
          id: {
            description: "id of the store",
            type: "number",
            example: 5,
          },
          name: {
            description: "Name of the store",
            type: "string",
            example: "Apple Store",
          },
          city: {
            description: "City where the store is located",
            type: "string",
            example: "Hanoi",
          },
          country: {
            description: "Country where the store is located",
            type: "string",
            example: "Vietnam",
          },
          street_name: {
            description: "Street name of the store's address",
            type: "string",
            example: "Le Loi Street",
          },
          street_number: {
            description: "Street number of the store's address",
            type: "string",
            example: "123",
          },
          manager: {
            description: "Name of the store manager",
            type: "string",
            example: "Nguyen Van A",
          },
        },
      },

      Price: {
        required: ["product_id", "price_change", "change_status", "new_price"],
        properties: {
          id: {
            description: "id of the price",
            type: "number",
            example: 5,
          },
          product_id: {
            description: "Product id",
            type: "integer",
            example: 101,
          },
          price_change: {
            description: "Amount of price change for the product",
            type: "number",
            format: "float",
            example: 5.99,
          },
          change_status: {
            description: "Indicates whether the price is increasing or decreasing",
            type: "string",
            example: "+",
          },
          new_price: {
            description: "Price after change",
            type: "number",
            format: "float",
            example: 5.99,
          },
        },
      },

      Campaign: {
        required: ["name", "time_start", "time_end", "description", "status"],
        properties: {
          id: {
            description: "id of the campaign",
            type: "number",
            example: 5,
          },
          name: {
            description: "Name of the campaign",
            type: "string",
            example: "Holiday Sales",
          },
          time_start: {
            description: "Start time of the campaign",
            type: "string",
            example: "2023-09-01 ",
          },
          time_end: {
            description: "End time of the campaign",
            type: "string",
            example: "2023-12-31",
          },
          description: {
            description: "Description of the campaign",
            type: "string",
            example: "Special sales for the holiday season.",
          },
          status: {
            description: "Status of the campaign",
            type: "string",
            example: "Active",
          },
        },
      },

      Customer: {
        required: [
          "name",
          "gender",
          "email",
          "phone",
          "city",
          "country",
          "street_name",
          "street_number",
          "market_segment",
        ],
        properties: {
          id: {
            description: "id of the customer",
            type: "number",
            example: 5,
          },
          name: {
            description: "Name of the customer",
            type: "string",
            example: "John Doe",
          },
          gender: {
            description: "Gender of the customer",
            type: "string",
            example: "Male",
          },
          email: {
            description: "Email address of the customer",
            type: "string",
            example: "johndoe@example.com",
          },
          phone: {
            description: "Phone number of the customer",
            type: "string",
            example: "+1234567890",
          },
          city: {
            description: "City of the customer's residence",
            type: "string",
            example: "New York",
          },
          country: {
            description: "Country of the customer's residence",
            type: "string",
            example: "USA",
          },
          street_name: {
            description: "Street name of the customer's address",
            type: "string",
            example: "Broadway",
          },
          street_number: {
            description: "Street number of the customer's address",
            type: "string",
            example: "123",
          },
          market_segment: {
            description: "Market segment the customer belongs to",
            type: "string",
            example: "Retail",
          },
        },
      },

      DateOrder: {
        required: ["day_number_of_month", "month_number_of_year", "year"],
        properties: {
          id: {
            description: "id of the date order",
            type: "number",
            example: 5,
          },
          day_number_of_month: {
            description: "Day number within the month",
            type: "integer",
            example: 12,
          },
          month_number_of_year: {
            description: "Month number within the year",
            type: "integer",
            example: 8,
          },
          year: {
            description: "Year of the order date",
            type: "integer",
            example: 2023,
          },
        },
      },

      Delivery: {
        required: ["delivery_method", "price"],
        properties: {
          id: {
            description: "id of the delivery",
            type: "number",
            example: 5,
          },
          delivery_method: {
            description: "Method of delivery",
            type: "string",
            example: "Courier",
          },
          price: {
            description: "Price for the delivery method",
            type: "number",
            format: "float",
            example: 15.5,
          },
        },
      },

      Discount: {
        required: [
          "id_campaign",
          "id_strategy",
          "id_discount_type",
          "id_discount_value",
        ],
        properties: {
          id: {
            description: "id of the discount",
            type: "number",
            example: 5,
          },
          id_campaign: {
            description: "Identifier for the campaign",
            type: "integer",
            example: 101,
          },
          id_strategy: {
            description: "Identifier for the strategy",
            type: "integer",
            example: 202,
          },
          id_discount_type: {
            description: "Identifier for the discount type",
            type: "integer",
            example: 303,
          },
          id_discount_value: {
            description: "Identifier for the discount value",
            type: "integer",
            example: 404,
          },
        },
      },

      DiscountType: {
        required: [
          "quality_discount",
          "invoice_value_condition",
          "quantity_condition",
          "brand",
          "product_id",
          "description",
          "is_delete",
        ],
        properties: {
          id: {
            description: "id of the discount type",
            type: "number",
            example: 5,
          },
          quality_discount: {
            description: "Quality discount percentage or value",
            type: "integer",
            example: 10,
          },
          invoice_value_condition: {
            description: "Invoice value condition for applying the discount",
            type: "number",
            format: "float",
            example: 150.5,
          },
          quantity_condition: {
            description: "Quantity condition for applying the discount",
            type: "integer",
            example: 3,
          },
          brand: {
            description: "Brand related to the discount",
            type: "string",
            example: "BrandName",
          },
          product_id: {
            description: "Product identifier related to the discount",
            type: "integer",
            example: 102,
          },
          description: {
            description: "Description of the discount type",
            type: "string",
            example: "Discount for buying in bulk.",
          },
          is_delete: {
            description: "Indicates if the discount type is deleted",
            type: "boolean",
            example: false,
          },
        },
      },

      DiscountValue: {
        required: ["display_name", "value", "max_value"],
        properties: {
          id: {
            description: "id of the discount value",
            type: "number",
            example: 5,
          },
          display_name: {
            description: "Display name for the discount value",
            type: "string",
            example: "Seasonal Discount",
          },
          value: {
            description: "Discount value, could represent a percentage or a flat discount",
            type: "number",
            format: "float",
            example: 15.5,
          },
          max_value: {
            description: "Maximum value for the discount if applicable",
            type: "number",
            format: "float",
            example: 50.0,
          },
        },
      },

      Employee: {
        required: [
          "name",
          "city",
          "country",
          "street_name",
          "street_number",
          "gender",
          "age",
          "salary_level",
          "status",
          "education_level",
          "last_rating",
        ],
        properties: {
          id: {
            description: "id of the employee",
            type: "number",
            example: 5,
          },
          name: {
            description: "Name of the employee",
            type: "string",
            example: "John Doe",
          },
          city: {
            description: "City where the employee resides",
            type: "string",
            example: "Los Angeles",
          },
          country: {
            description: "Country where the employee resides",
            type: "string",
            example: "USA",
          },
          street_name: {
            description: "Street name of the employee's address",
            type: "string",
            example: "Main Street",
          },
          street_number: {
            description: "Street number of the employee's address",
            type: "string",
            example: "123A",
          },
          gender: {
            description: "Gender of the employee",
            type: "string",
            enum: ["Male", "Female", "Other"],
            example: "Male",
          },
          age: {
            description: "Age of the employee",
            type: "integer",
            example: 30,
          },
          salary_level: {
            description: "Salary level or band of the employee",
            type: "string",
            example: "S2",
          },
          status: {
            description: "Employment status of the employee",
            type: "string",
            example: "Active",
          },
          education_level: {
            description: "Education level of the employee",
            type: "string",
            example: "Bachelor's Degree",
          },
          last_rating: {
            description: "Last performance rating of the employee",
            type: "number",
            format: "float",
            example: 4.5,
          },
        },
      },

      Order: {
        required: [
          "customer_id",
          "product_id",
          "store_id",
          "employee_id",
          "payment_id",
          "delivery_id",
          "date_order_id",
          "discount_id",
          "total_amount_order",
          "discount_amount",
        ],
        properties: {
          id: {
            description: "id of the order",
            type: "number",
            example: 5,
          },
          customer_id: {
            description: "ID of the customer placing the order",
            type: "integer",
            example: 1,
          },
          product_id: {
            description: "ID of the product placing the order",
            type: "integer",
            example: 201,
          },
          store_id: {
            description: "ID of the store where the order was placed",
            type: "integer",
            example: 15,
          },
          employee_id: {
            description: "ID of the employee handling the order",
            type: "integer",
            example: 45,
          },
          payment_id: {
            description: "ID associated with the payment for the order",
            type: "integer",
            example: 1245,
          },
          delivery_id: {
            description: "ID related to the delivery details of the order",
            type: "integer",
            example: 301,
          },
          date_order_id: {
            description: "ID corresponding to the date of the order",
            type: "integer",
            example: 9201,
          },
          discount_id: {
            description: "ID related to any discount applied to the order",
            type: "integer",
            example: 7,
          },
          total_amount_order: {
            description: "Total amount for the order",
            type: "number",
            format: "float",
            example: 99.99,
          },
          discount_amount: {
            description: "Amount discounted from the total order price",
            type: "number",
            format: "float",
            example: 10.0,
          },
        },
      },

      Payment: {
        required: ["payment_method"],
        properties: {
          id: {
            description: "id of the payment",
            type: "number",
            example: 5,
          },
          payment_method: {
            description: "Method used for the payment",
            type: "string",
            example: "Credit Card",
          },
        },
      },

      Strategy: {
        required: ["name", "time_start", "time_end", "description", "status"],
        properties: {
          id: {
            description: "id of the strategy",
            type: "number",
            example: 5,
          },
          name: {
            description: "Name of the strategy",
            type: "string",
            example: "Winter Sale",
          },
          time_start: {
            description: "Start time of the strategy",
            type: "string",
            example: "2023-01-01",
          },
          time_end: {
            description: "End time of the strategy",
            type: "string",
            example: "2023-03-01",
          },
          description: {
            description: "Description of the strategy",
            type: "string",
            example: "Strategy for winter sales campaign.",
          },
          status: {
            description: "Status of the strategy",
            type: "string",
            example: "Active",
          },
        },
      },

      Error: {
        description: "<table>\n  <tr>\n    <th>Code</th>\n    <th>Description</th>\n  </tr>\n  <tr>\n    <td>illegal_input</td>\n    <td>The input is invalid.</td>\n  </tr>\n  <tr>\n    <td>not_found</td>\n    <td>The resource is not found.</td>\n  </tr>\n</table>\n",
        required: ["code", "message"],
        properties: {
          code: {
            type: "string",
            example: "illegal_input",
          },
        },
      },
    },
    parameters: {
      Id: {
        name: "id",
        in: "path",
        description: "ID of object",
        type: "integer",
        format: "int64",
        required: true,
        schema: {
          $ref: "#/components/parameters/Id",
        },
      },
      ProductNoID: {
        required: [
          "name",
          "description",
          "price",
          "stock",
          "unit",
          "category",
          "brand",
          "sale_status",
          "sales_type",
        ],
        properties: {
          name: {
            description: "Name of the product",
            type: "string",
            example: "Apple",
          },
          description: {
            description: "Description of the product",
            type: "string",
            example: "This is fresh food.",
          },
          price: {
            description: "Price of the product",
            type: "number",
            format: "float",
            example: 19.99,
          },
          stock: {
            description: "Stock quantity of the product",
            type: "integer",
            example: 100,
          },
          unit: {
            description: "Unit of product volume",
            type: "string",
            example: "kg",
          },
          category: {
            description: "Category of the product",
            type: "string",
            example: "Fruit",
          },
          brand: {
            description: "Brand of the product",
            type: "string",
            example: "NewZealand",
          },
          sale_status: {
            description: "Product sales status",
            type: "string",
            example: "In Stock",
          },
          sales_type: {
            description: "Sales type of the product",
            type: "string",
            example: "Online",
          },
        },
      },
      StoreNoID: {
        required: [
          "name",
          "city",
          "country",
          "street_name",
          "street_number",
          "manager",
        ],
        properties: {
          name: {
            description: "Name of the store",
            type: "string",
            example: "Apple Store",
          },
          city: {
            description: "City where the store is located",
            type: "string",
            example: "Hanoi",
          },
          country: {
            description: "Country where the store is located",
            type: "string",
            example: "Vietnam",
          },
          street_name: {
            description: "Street name of the store's address",
            type: "string",
            example: "Le Loi Street",
          },
          street_number: {
            description: "Street number of the store's address",
            type: "string",
            example: "123",
          },
          manager: {
            description: "Name of the store manager",
            type: "string",
            example: "Nguyen Van A",
          },
        },
      },
      PriceNoID: {
        required: ["product_id", "price_change", "change_status", "new_price"],
        properties: {
          id: {
            description: "id of the price",
            type: "number",
            example: 5,
          },
          product_id: {
            description: "Product id",
            type: "integer",
            example: 101,
          },
          price_change: {
            description: "Amount of price change for the product",
            type: "number",
            format: "float",
            example: 5.99,
          },
          change_status: {
            description: "Indicates whether the price is increasing or decreasing",
            type: "string",
            example: "+",
          },
          new_price: {
            description: "Price after change",
            type: "number",
            format: "float",
            example: 5.99,
          },
        },
      },
      CampaignNoID: {
        required: ["name", "time_start", "time_end", "description", "status"],
        properties: {
          name: {
            description: "Name of the campaign",
            type: "string",
            example: "Holiday Sales",
          },
          time_start: {
            description: "Start time of the campaign",
            type: "string",
            example: "2023-09-01 ",
          },
          time_end: {
            description: "End time of the campaign",
            type: "string",
            example: "2023-12-31",
          },
          description: {
            description: "Description of the campaign",
            type: "string",
            example: "Special sales for the holiday season.",
          },
          status: {
            description: "Status of the campaign",
            type: "string",
            example: "Active",
          },
        },
      },
      CustomerNoID: {
        required: [
          "name",
          "gender",
          "email",
          "phone",
          "city",
          "country",
          "street_name",
          "street_number",
          "market_segment",
        ],
        properties: {
          name: {
            description: "Name of the customer",
            type: "string",
            example: "John Doe",
          },
          gender: {
            description: "Gender of the customer",
            type: "string",
            example: "Male",
          },
          email: {
            description: "Email address of the customer",
            type: "string",
            example: "johndoe@example.com",
          },
          phone: {
            description: "Phone number of the customer",
            type: "string",
            example: "+1234567890",
          },
          city: {
            description: "City of the customer's residence",
            type: "string",
            example: "New York",
          },
          country: {
            description: "Country of the customer's residence",
            type: "string",
            example: "USA",
          },
          street_name: {
            description: "Street name of the customer's address",
            type: "string",
            example: "Broadway",
          },
          street_number: {
            description: "Street number of the customer's address",
            type: "string",
            example: "123",
          },
          market_segment: {
            description: "Market segment the customer belongs to",
            type: "string",
            example: "Retail",
          },
        },
      },
      DateOrderNoID: {
        required: ["day_number_of_month", "month_number_of_year", "year"],
        properties: {
          day_number_of_month: {
            description: "Day number within the month",
            type: "integer",
            example: 12,
          },
          month_number_of_year: {
            description: "Month number within the year",
            type: "integer",
            example: 8,
          },
          year: {
            description: "Year of the order date",
            type: "integer",
            example: 2023,
          },
        },
      },
      DeliveryNoID: {
        required: ["delivery_method", "price"],
        properties: {
          delivery_method: {
            description: "Method of delivery",
            type: "string",
            example: "Courier",
          },
          price: {
            description: "Price for the delivery method",
            type: "number",
            format: "float",
            example: 15.5,
          },
        },
      },
      DiscountNoID: {
        required: [
          "id_campaign",
          "id_strategy",
          "id_discount_type",
          "id_discount_value",
        ],
        properties: {
          id_campaign: {
            description: "Identifier for the campaign",
            type: "integer",
            example: 101,
          },
          id_strategy: {
            description: "Identifier for the strategy",
            type: "integer",
            example: 202,
          },
          id_discount_type: {
            description: "Identifier for the discount type",
            type: "integer",
            example: 303,
          },
          id_discount_value: {
            description: "Identifier for the discount value",
            type: "integer",
            example: 404,
          },
        },
      },
      DiscountTypeNoID: {
        required: [
          "quality_discount",
          "invoice_value_condition",
          "quantity_condition",
          "brand",
          "product_id",
          "description",
          "is_delete",
        ],
        properties: {
          quality_discount: {
            description: "Quality discount percentage or value",
            type: "integer",
            example: 10,
          },
          invoice_value_condition: {
            description: "Invoice value condition for applying the discount",
            type: "number",
            format: "float",
            example: 150.5,
          },
          quantity_condition: {
            description: "Quantity condition for applying the discount",
            type: "integer",
            example: 3,
          },
          brand: {
            description: "Brand related to the discount",
            type: "string",
            example: "BrandName",
          },
          product_id: {
            description: "Product identifier related to the discount",
            type: "integer",
            example: 102,
          },
          description: {
            description: "Description of the discount type",
            type: "string",
            example: "Discount for buying in bulk.",
          },
          is_delete: {
            description: "Indicates if the discount type is deleted",
            type: "boolean",
            example: false,
          },
        },
      },
      DiscountValueNoID: {
        required: ["display_name", "value", "max_value"],
        properties: {
          display_name: {
            description: "Display name for the discount value",
            type: "string",
            example: "Seasonal Discount",
          },
          value: {
            description: "Discount value, could represent a percentage or a flat discount",
            type: "number",
            format: "float",
            example: 15.5,
          },
          max_value: {
            description: "Maximum value for the discount if applicable",
            type: "number",
            format: "float",
            example: 50.0,
          },
        },
      },
      EmployeeNoID: {
        required: [
          "name",
          "city",
          "country",
          "street_name",
          "street_number",
          "gender",
          "age",
          "salary_level",
          "status",
          "education_level",
          "last_rating",
        ],
        properties: {
          name: {
            description: "Name of the employee",
            type: "string",
            example: "John Doe",
          },
          city: {
            description: "City where the employee resides",
            type: "string",
            example: "Los Angeles",
          },
          country: {
            description: "Country where the employee resides",
            type: "string",
            example: "USA",
          },
          street_name: {
            description: "Street name of the employee's address",
            type: "string",
            example: "Main Street",
          },
          street_number: {
            description: "Street number of the employee's address",
            type: "string",
            example: "123A",
          },
          gender: {
            description: "Gender of the employee",
            type: "string",
            enum: ["Male", "Female", "Other"],
            example: "Male",
          },
          age: {
            description: "Age of the employee",
            type: "integer",
            example: 30,
          },
          salary_level: {
            description: "Salary level or band of the employee",
            type: "string",
            example: "S2",
          },
          status: {
            description: "Employment status of the employee",
            type: "string",
            example: "Active",
          },
          education_level: {
            description: "Education level of the employee",
            type: "string",
            example: "Bachelor's Degree",
          },
          last_rating: {
            description: "Last performance rating of the employee",
            type: "number",
            format: "float",
            example: 4.5,
          },
        },
      },
      OrderNoID: {
        required: [
          "customer_id",
          "product_id",
          "store_id",
          "employee_id",
          "payment_id",
          "delivery_id",
          "date_order_id",
          "discount_id",
          "total_amount_order",
          "discount_amount",
        ],
        properties: {
          customer_id: {
            description: "ID of the customer placing the order",
            type: "integer",
            example: 1,
          },
          product_id: {
            description: "ID of the product",
            type: "integer",
            example: 201,
          },
          store_id: {
            description: "ID of the store where the order was placed",
            type: "integer",
            example: 15,
          },
          employee_id: {
            description: "ID of the employee handling the order",
            type: "integer",
            example: 45,
          },
          payment_id: {
            description: "ID associated with the payment for the order",
            type: "integer",
            example: 1245,
          },
          delivery_id: {
            description: "ID related to the delivery details of the order",
            type: "integer",
            example: 301,
          },
          date_order_id: {
            description: "ID corresponding to the date of the order",
            type: "integer",
            example: 9201,
          },
          discount_id: {
            description: "ID related to any discount applied to the order",
            type: "integer",
            example: 7,
          },
          total_amount_order: {
            description: "Total amount for the order",
            type: "number",
            format: "float",
            example: 99.99,
          },
          discount_amount: {
            description: "Amount discounted from the total order price",
            type: "number",
            format: "float",
            example: 10.0,
          },
        },
      },
      PaymentNoID: {
        required: ["payment_method"],
        properties: {
          payment_method: {
            description: "Method used for the payment",
            type: "string",
            example: "Credit Card",
          },
        },
      },
      StrategyNoID: {
        required: ["name", "time_start", "time_end", "description", "status"],
        properties: {
          name: {
            description: "Name of the strategy",
            type: "string",
            example: "Winter Sale",
          },
          time_start: {
            description: "Start time of the strategy",
            type: "string",
            example: "2023-01-01",
          },
          time_end: {
            description: "End time of the strategy",
            type: "string",
            example: "2023-03-01",
          },
          description: {
            description: "Description of the strategy",
            type: "string",
            example: "Strategy for winter sales campaign.",
          },
          status: {
            description: "Status of the strategy",
            type: "string",
            example: "Active",
          },
        },
      },
      Name: {
        name: "name",
        in: "query",
        description: "Name object want to search",
        type: "string",
        format: "string",
        required: false,
        schema: {
          $ref: "#/components/parameters/Name",
        },
      },
      Description: {
        name: "description",
        in: "query",
        description: "Description object want to search",
        type: "string",
        format: "string",
        required: false,
        schema: {
          $ref: "#/components/parameters/Description",
        },
      },
      Method: {
        name: "method",
        in: "query",
        description: "Method want to search",
        type: "string",
        format: "string",
        required: true,
        schema: {
          $ref: "#/components/parameters/Method",
        },
      },
      ProductID: {
        name: "productid",
        in: "query",
        description: "id product want to search",
        type: "number",
        format: "int",
        required: true,
        schema: {
          $ref: "#/components/parameters/ProductID",
        },
      },
      Product: {
        name: "Product",
        in: "path",
        description: "Resource Product",
        required: true,
        schema: {
          $ref: "#/components/schemas/Product",
        },
      },
      Store: {
        name: "Store",
        in: "path",
        description: "Resource Store",
        required: true,
        schema: {
          $ref: "#/components/schemas/Store",
        },
      },
      Price: {
        name: "Price",
        in: "path",
        description: "Resource Price",
        required: true,
        schema: {
          $ref: "#/components/schemas/Price",
        },
      },
      Campaign: {
        name: "Campaign",
        in: "path",
        description: "Resource Campaign",
        required: true,
        schema: {
          $ref: "#/components/schemas/Campaign",
        },
      },
      Customer: {
        name: "Customer",
        in: "path",
        description: "Resource Customer",
        required: true,
        schema: {
          $ref: "#/components/schemas/Customer",
        },
      },
      DateOrder: {
        name: "DateOrder",
        in: "path",
        description: "Resource DateOrder",
        required: true,
        schema: {
          $ref: "#/components/schemas/DateOrder",
        },
      },
      Delivery: {
        name: "Delivery",
        in: "path",
        description: "Resource Delivery",
        required: true,
        schema: {
          $ref: "#/components/schemas/Delivery",
        },
      },
      Discount: {
        name: "Discount",
        in: "path",
        description: "Resource Discount",
        required: true,
        schema: {
          $ref: "#/components/schemas/Discount",
        },
      },
      DiscountType: {
        name: "DiscountType",
        in: "path",
        description: "Resource DiscountType",
        required: true,
        schema: {
          $ref: "#/components/schemas/DiscountType",
        },
      },
      DiscountValue: {
        name: "DiscountValue",
        in: "path",
        description: "Resource DiscountValue",
        required: true,
        schema: {
          $ref: "#/components/schemas/DiscountValue",
        },
      },
      Employee: {
        name: "Employee",
        in: "path",
        description: "Resource Employee",
        required: true,
        schema: {
          $ref: "#/components/schemas/Employee",
        },
      },
      Order: {
        name: "Order",
        in: "path",
        description: "Resource Order",
        required: true,
        schema: {
          $ref: "#/components/schemas/Order",
        },
      },
      Payment: {
        name: "Payment",
        in: "path",
        description: "Resource Payment",
        required: true,
        schema: {
          $ref: "#/components/schemas/Payment",
        },
      },
      Strategy: {
        name: "Strategy",
        in: "path",
        description: "Resource Strategy",
        required: true,
        schema: {
          $ref: "#/components/schemas/Strategy",
        },
      },
      Success: {
        ProductResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                product: {
                  type: "object",
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    name: {
                      type: "string",
                      example: "Apple",
                    },
                    description: {
                      type: "string",
                      example: "Apple is a fresh food",
                    },
                    price: {
                      type: "number",
                      example: 19.9,
                    },
                    stock: {
                      type: "integer",
                      example: 100,
                    },
                    unit: {
                      type: "string",
                      example: "kg",
                    },
                    category: {
                      type: "string",
                      example: "Fruit",
                    },
                    brand: {
                      type: "string",
                      example: "NewZealand",
                    },
                    sale_status: {
                      type: "string",
                      example: "availible",
                    },
                    sales_type: {
                      type: "string",
                      example: "offline",
                    },
                  },
                  required: [
                    "name",
                    "description",
                    "price",
                    "stock",
                    "unit",
                    "category",
                    "brand",
                    "sale_status",
                    "sales_type",
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["product", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListProductResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_product: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    properties: {
                      id: {
                        type: "number",
                      },
                      name: {
                        type: "string",
                      },
                      description: {
                        type: "string",
                      },
                      price: {
                        type: "number",
                      },
                      stock: {
                        type: "integer",
                      },
                      unit: {
                        type: "string",
                      },
                      category: {
                        type: "string",
                      },
                      brand: {
                        type: "string",
                      },
                      sale_status: {
                        type: "string",
                      },
                      sales_type: {
                        type: "string",
                      },
                    },
                    required: [
                      "name",
                      "description",
                      "price",
                      "stock",
                      "unit",
                      "category",
                      "brand",
                      "sale_status",
                      "sales_type",
                    ],
                  },
                  example: [{
                      id: 1,
                      name: "Product 1",
                      description: "Description for Product 1",
                      price: 100.0,
                      stock: 50,
                      unit: "kg",
                      category: "Clothing",
                      brand: "BrandA",
                      sale_status: "In Stock",
                      sales_type: "Online",
                    },
                    {
                      id: 2,
                      name: "Product 2",
                      description: "Description for Product 2",
                      price: 110.0,
                      stock: 40,
                      unit: "L",
                      category: "Electronics",
                      brand: "BrandB",
                      sale_status: "Out of Stock",
                      sales_type: "Offline",
                    },
                    {
                      id: 3,
                      name: "Product 3",
                      description: "Description for Product 3",
                      price: 110.0,
                      stock: 40,
                      unit: "L",
                      category: "Electronics",
                      brand: "BrandB",
                      sale_status: "Out of Stock",
                      sales_type: "Offline",
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["products", "isError", "status"],
            },
          },
          required: ["data"],
        },
        StoreResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                store: {
                  type: "object",
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    name: {
                      type: "string",
                      example: " Store",
                    },
                    city: {
                      type: "string",
                      example: "HaNoi",
                    },
                    country: {
                      type: "string",
                      example: "VietNam",
                    },
                    street_name: {
                      type: "string",
                      example: "LeLoiStreet",
                    },
                    street_number: {
                      type: "string",
                      example: "1060/2",
                    },
                    manager: {
                      type: "string",
                      example: "NguyenVanA",
                    },
                  },
                  required: [
                    "name",
                    "city",
                    "country",
                    "street_name",
                    "street_number",
                    "manager",
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["store", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListStoreResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_store: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      name: {
                        type: "string",
                        example: " Store",
                      },
                      city: {
                        type: "string",
                        example: "HaNoi",
                      },
                      country: {
                        type: "string",
                        example: "VietNam",
                      },
                      street_name: {
                        type: "string",
                        example: "LeLoiStreet",
                      },
                      street_number: {
                        type: "string",
                        example: "1060/2",
                      },
                      manager: {
                        type: "string",
                        example: "NguyenVanA",
                      },
                    },
                    required: [
                      "name",
                      "city",
                      "country",
                      "street_name",
                      "street_number",
                      "manager",
                    ],
                  },
                  example: [{
                      id: 1,
                      name: "Store 1",
                      city: "Ho Chi Minh ",
                      country: "VietNam",
                      street_name: "LeLoi",
                      street_number: "10/2",
                      manager: "NguyenVanA",
                    },
                    {
                      id: 2,
                      name: "Store 2",
                      city: "Ho Chi Minh ",
                      country: "VietNam",
                      street_name: "LeLoi",
                      street_number: "10/2",
                      manager: "NguyenVanB",
                    },
                    {
                      id: 3,
                      name: "Store 3",
                      city: "Ho Chi Minh ",
                      country: "VietNam",
                      street_name: "LeLoi",
                      street_number: "10/2",
                      manager: "NguyenVanC",
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["stores", "isError", "status"],
            },
          },
          required: ["data"],
        },
        PriceResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                price: {
                  type: "object",
                  required: [
                    "product_id",
                    "price_change",
                    "change_status",
                    "new_price",
                  ],
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    product_id: {
                      description: "Product id",
                      type: "integer",
                      example: 101,
                    },
                    price_change: {
                      description: "Amount of price change for the product",
                      type: "number",
                      format: "float",
                      example: 5.99,
                    },
                    change_status: {
                      description: "Indicates whether the price is increasing or decreasing",
                      type: "string",
                      example: "+",
                    },
                    new_price: {
                      description: "Price after change",
                      type: "number",
                      format: "float",
                      example: 5.99,
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["price", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListPriceResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_price: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: [
                      "product_id",
                      "price_change",
                      "change_status",
                      "new_price",
                    ],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      product_id: {
                        description: "Product id",
                        type: "integer",
                        example: 101,
                      },
                      price_change: {
                        description: "Amount of price change for the product",
                        type: "number",
                        format: "float",
                        example: 5.99,
                      },
                      change_status: {
                        description: "Indicates whether the price is increasing or decreasing",
                        type: "string",
                        example: "+",
                      },
                      new_price: {
                        description: "Price after change",
                        type: "number",
                        format: "float",
                        example: 5.99,
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      product_id: 1,
                      price_change: 5000,
                      change_status: "+",
                      new_price: 15000,
                    },
                    {
                      id: 2,
                      product_id: 2,
                      price_change: 10000,
                      change_status: "",
                      new_price: 15000,
                    },
                    {
                      id: 3,
                      product_id: 3,
                      price_change: 5000,
                      change_status: "+",
                      new_price: 15000,
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["prices", "isError", "status"],
            },
          },
          required: ["data"],
        },
        CampaignResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                campaign: {
                  type: "object",
                  required: [
                    "name",
                    "time_start",
                    "time_end",
                    "description",
                    "status",
                  ],
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    name: {
                      description: "Name of the campaign",
                      type: "string",
                      example: "Holiday Sales",
                    },
                    time_start: {
                      description: "Start time of the campaign",
                      type: "string",
                      example: "2023-09-01 ",
                    },
                    time_end: {
                      description: "End time of the campaign",
                      type: "string",
                      example: "2023-12-31",
                    },
                    description: {
                      description: "Description of the campaign",
                      type: "string",
                      example: "Special sales for the holiday season.",
                    },
                    status: {
                      description: "Status of the campaign",
                      type: "string",
                      example: "Active",
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["campaign", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListCampaignResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_campaign: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: [
                      "name",
                      "time_start",
                      "time_end",
                      "description",
                      "status",
                    ],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      name: {
                        description: "Name of the campaign",
                        type: "string",
                        example: "Holiday Sales",
                      },
                      time_start: {
                        description: "Start time of the campaign",
                        type: "string",
                        example: "2023-09-01 ",
                      },
                      time_end: {
                        description: "End time of the campaign",
                        type: "string",
                        example: "2023-12-31",
                      },
                      description: {
                        description: "Description of the campaign",
                        type: "string",
                        example: "Special sales for the holiday season.",
                      },
                      status: {
                        description: "Status of the campaign",
                        type: "string",
                        example: "Active",
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      name: "Winter Sales Campaign",
                      time_start: "2023-12-01",
                      time_end: "2023-12-31",
                      description: "Campaign for winter sales, aiming to increase end-of-year sales.",
                      status: "Planned",
                    },
                    {
                      id: 2,
                      name: "Summer Sales Campaign",
                      time_start: "2023-12-01",
                      time_end: "2023-12-31",
                      description: "Campaign for Summer sales, aiming to increase end-of-year sales.",
                      status: "Planned",
                    },
                    {
                      id: 3,
                      name: "Sunshine Sales Campaign",
                      time_start: "2023-12-01",
                      time_end: "2023-12-31",
                      description: "Campaign for Sunshine sales, aiming to increase end-of-year sales.",
                      status: "Planned",
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["campaigns", "isError", "status"],
            },
          },
          required: ["data"],
        },
        CustomerResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                customer: {
                  type: "object",
                  required: [
                    "name",
                    "gender",
                    "email",
                    "phone",
                    "city",
                    "country",
                    "street_name",
                    "street_number",
                    "market_segment",
                  ],
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    name: {
                      description: "Name of the customer",
                      type: "string",
                      example: "John Doe",
                    },
                    gender: {
                      description: "Gender of the customer",
                      type: "string",
                      example: "Male",
                    },
                    email: {
                      description: "Email address of the customer",
                      type: "string",
                      example: "johndoe@example.com",
                    },
                    phone: {
                      description: "Phone number of the customer",
                      type: "string",
                      example: "+1234567890",
                    },
                    city: {
                      description: "City of the customer's residence",
                      type: "string",
                      example: "New York",
                    },
                    country: {
                      description: "Country of the customer's residence",
                      type: "string",
                      example: "USA",
                    },
                    street_name: {
                      description: "Street name of the customer's address",
                      type: "string",
                      example: "Broadway",
                    },
                    street_number: {
                      description: "Street number of the customer's address",
                      type: "string",
                      example: "123",
                    },
                    market_segment: {
                      description: "Market segment the customer belongs to",
                      type: "string",
                      example: "Retail",
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["customer", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListCustomerResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_customer: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: [
                      "name",
                      "gender",
                      "email",
                      "phone",
                      "city",
                      "country",
                      "street_name",
                      "street_number",
                      "market_segment",
                    ],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      name: {
                        description: "Name of the customer",
                        type: "string",
                        example: "John Doe",
                      },
                      gender: {
                        description: "Gender of the customer",
                        type: "string",
                        example: "Male",
                      },
                      email: {
                        description: "Email address of the customer",
                        type: "string",
                        example: "johndoe@example.com",
                      },
                      phone: {
                        description: "Phone number of the customer",
                        type: "string",
                        example: "+1234567890",
                      },
                      city: {
                        description: "City of the customer's residence",
                        type: "string",
                        example: "New York",
                      },
                      country: {
                        description: "Country of the customer's residence",
                        type: "string",
                        example: "USA",
                      },
                      street_name: {
                        description: "Street name of the customer's address",
                        type: "string",
                        example: "Broadway",
                      },
                      street_number: {
                        description: "Street number of the customer's address",
                        type: "string",
                        example: "123",
                      },
                      market_segment: {
                        description: "Market segment the customer belongs to",
                        type: "string",
                        example: "Retail",
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      name: "John Doe",
                      gender: "Male",
                      email: "johndoe@example.com",
                      phone: "+1234567890",
                      city: "New York",
                      country: "USA",
                      street_name: "Broadway",
                      street_number: "123A",
                      market_segment: "Electronics",
                    },
                    {
                      id: 2,
                      name: "John Doe",
                      gender: "Male",
                      email: "johndoe@example.com",
                      phone: "+34534534534",
                      city: "New York",
                      country: "USA",
                      street_name: "Broadway",
                      street_number: "123A",
                      market_segment: "Electronics",
                    },
                    {
                      id: 3,
                      name: "John Doe",
                      gender: "Male",
                      email: "johndoe@example.com",
                      phone: "+678678678",
                      city: "New York",
                      country: "USA",
                      street_name: "Broadway",
                      street_number: "123A",
                      market_segment: "Electronics",
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["customers", "isError", "status"],
            },
          },
          required: ["data"],
        },
        DateOrderResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                id: {
                  type: "number",
                  example: 5,
                },
                date_order: {
                  type: "object",
                  required: [
                    "day_number_of_month",
                    "month_number_of_year",
                    "year",
                  ],
                  properties: {
                    day_number_of_month: {
                      description: "Day number within the month",
                      type: "integer",
                      example: 12,
                    },
                    month_number_of_year: {
                      description: "Month number within the year",
                      type: "integer",
                      example: 8,
                    },
                    year: {
                      description: "Year of the order date",
                      type: "integer",
                      example: 2023,
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["date_order", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListDateOrderResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_date_order: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: [
                      "day_number_of_month",
                      "month_number_of_year",
                      "year",
                    ],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      day_number_of_month: {
                        description: "Day number within the month",
                        type: "integer",
                        example: 12,
                      },
                      month_number_of_year: {
                        description: "Month number within the year",
                        type: "integer",
                        example: 8,
                      },
                      year: {
                        description: "Year of the order date",
                        type: "integer",
                        example: 2023,
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      day_number_of_month: 15,
                      month_number_of_year: 5,
                      year: 2023,
                    },
                    {
                      id: 2,
                      day_number_of_month: 15,
                      month_number_of_year: 9,
                      year: 2023,
                    },
                    {
                      id: 3,
                      day_number_of_month: 15,
                      month_number_of_year: 7,
                      year: 2023,
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["date_orders", "isError", "status"],
            },
          },
          required: ["data"],
        },
        DeliveryResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                delivery: {
                  type: "object",
                  required: ["delivery_method", "price"],
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    delivery_method: {
                      description: "Method of delivery",
                      type: "string",
                      example: "Courier",
                    },
                    price: {
                      description: "Price for the delivery method",
                      type: "number",
                      format: "float",
                      example: 15.5,
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["delivery", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListDeliveryResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_delivery: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: ["delivery_method", "price"],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      delivery_method: {
                        description: "Method of delivery",
                        type: "string",
                        example: "Courier",
                      },
                      price: {
                        description: "Price for the delivery method",
                        type: "number",
                        format: "float",
                        example: 15.5,
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      delivery_method: "Delivery in 72h",
                      price: 45,
                    },
                    {
                      id: 2,
                      delivery_method: "Delivery in 24h",
                      price: 90,
                    },
                    {
                      id: 3,
                      delivery_method: "Delivery in 12h",
                      price: 180,
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["deliveries", "isError", "status"],
            },
          },
          required: ["data"],
        },
        DiscountResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                discount: {
                  type: "object",
                  required: [
                    "id_campaign",
                    "id_strategy",
                    "id_discount_type",
                    "id_discount_value",
                  ],
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    id_campaign: {
                      description: "Identifier for the campaign",
                      type: "integer",
                      example: 101,
                    },
                    id_strategy: {
                      description: "Identifier for the strategy",
                      type: "integer",
                      example: 202,
                    },
                    id_discount_type: {
                      description: "Identifier for the discount type",
                      type: "integer",
                      example: 303,
                    },
                    id_discount_value: {
                      description: "Identifier for the discount value",
                      type: "integer",
                      example: 404,
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["discount", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListDiscountResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_discount: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: [
                      "id_campaign",
                      "id_strategy",
                      "id_discount_type",
                      "id_discount_value",
                    ],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      id_campaign: {
                        description: "Identifier for the campaign",
                        type: "integer",
                        example: 101,
                      },
                      id_strategy: {
                        description: "Identifier for the strategy",
                        type: "integer",
                        example: 202,
                      },
                      id_discount_type: {
                        description: "Identifier for the discount type",
                        type: "integer",
                        example: 303,
                      },
                      id_discount_value: {
                        description: "Identifier for the discount value",
                        type: "integer",
                        example: 404,
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      id_campaign: 1,
                      id_strategy: 1,
                      id_discount_type: 1,
                      id_discount_value: 1,
                    },
                    {
                      id: 2,
                      id_campaign: 2,
                      id_strategy: 2,
                      id_discount_type: 2,
                      id_discount_value: 2,
                    },
                    {
                      id: 3,
                      id_campaign: 3,
                      id_strategy: 3,
                      id_discount_type: 3,
                      id_discount_value: 3,
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["discounts", "isError", "status"],
            },
          },
          required: ["data"],
        },
        DiscountTypeResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                discount_type: {
                  type: "object",
                  required: [
                    "quality_discount",
                    "invoice_value_condition",
                    "quantity_condition",
                    "brand",
                    "product_id",
                    "description",
                    "is_delete",
                  ],
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    quality_discount: {
                      description: "Quality discount percentage or value",
                      type: "integer",
                      example: 10,
                    },
                    invoice_value_condition: {
                      description: "Invoice value condition for applying the discount",
                      type: "number",
                      format: "float",
                      example: 150.5,
                    },
                    quantity_condition: {
                      description: "Quantity condition for applying the discount",
                      type: "integer",
                      example: 3,
                    },
                    brand: {
                      description: "Brand related to the discount",
                      type: "string",
                      example: "BrandName",
                    },
                    product_id: {
                      description: "Product identifier related to the discount",
                      type: "integer",
                      example: 102,
                    },
                    description: {
                      description: "Description of the discount type",
                      type: "string",
                      example: "Discount for buying in bulk.",
                    },
                    is_delete: {
                      description: "Indicates if the discount type is deleted",
                      type: "boolean",
                      example: false,
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["discount_type", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListDiscountTypeResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_discount_type: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: [
                      "quality_discount",
                      "invoice_value_condition",
                      "quantity_condition",
                      "brand",
                      "product_id",
                      "description",
                      "is_delete",
                    ],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      quality_discount: {
                        description: "Quality discount percentage or value",
                        type: "integer",
                        example: 10,
                      },
                      invoice_value_condition: {
                        description: "Invoice value condition for applying the discount",
                        type: "number",
                        format: "float",
                        example: 150.5,
                      },
                      quantity_condition: {
                        description: "Quantity condition for applying the discount",
                        type: "integer",
                        example: 3,
                      },
                      brand: {
                        description: "Brand related to the discount",
                        type: "string",
                        example: "BrandName",
                      },
                      product_id: {
                        description: "Product identifier related to the discount",
                        type: "integer",
                        example: 102,
                      },
                      description: {
                        description: "Description of the discount type",
                        type: "string",
                        example: "Discount for buying in bulk.",
                      },
                      is_delete: {
                        description: "Indicates if the discount type is deleted",
                        type: "boolean",
                        example: false,
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      quality_discount: 5,
                      invoice_value_condition: 500.5,
                      quantity_condition: 10,
                      brand: "Nike",
                      product_id: 12345,
                      description: "Giảm giá dành cho sản phẩm chất lượng cao",
                    },
                    {
                      id: 2,
                      quality_discount: 6,
                      invoice_value_condition: 500.5,
                      quantity_condition: 10,
                      brand: "Apple",
                      product_id: 12345,
                      description: "Giảm giá dành cho sản phẩm chất lượng cao",
                    },
                    {
                      id: 3,
                      quality_discount: 7,
                      invoice_value_condition: 500.5,
                      quantity_condition: 10,
                      brand: "Kiwi",
                      product_id: 12345,
                      description: "Giảm giá dành cho sản phẩm chất lượng cao",
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["discount_types", "isError", "status"],
            },
          },
          required: ["data"],
        },
        DiscountValueResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                discount_value: {
                  type: "object",
                  required: ["display_name", "value", "max_value"],
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    display_name: {
                      description: "Display name for the discount value",
                      type: "string",
                      example: "Seasonal Discount",
                    },
                    value: {
                      description: "Discount value, could represent a percentage or a flat discount",
                      type: "number",
                      format: "float",
                      example: 15.5,
                    },
                    max_value: {
                      description: "Maximum value for the discount if applicable",
                      type: "number",
                      format: "float",
                      example: 50.0,
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["discount_value", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListDiscountValueResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_discount_value: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: ["display_name", "value", "max_value"],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      display_name: {
                        description: "Display name for the discount value",
                        type: "string",
                        example: "Seasonal Discount",
                      },
                      value: {
                        description: "Discount value, could represent a percentage or a flat discount",
                        type: "number",
                        format: "float",
                        example: 15.5,
                      },
                      max_value: {
                        description: "Maximum value for the discount if applicable",
                        type: "number",
                        format: "float",
                        example: 50.0,
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      display_name: "Giảm giá đặc biệt",
                      value: 50.25,
                      max_value: 100.0,
                    },
                    {
                      id: 2,
                      display_name: "Giảm giá đặc biệt 1",
                      value: 50.25,
                      max_value: 100.0,
                    },
                    {
                      id: 3,
                      display_name: "Giảm giá đặc biệt 2",
                      value: 50.25,
                      max_value: 100.0,
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["discount_values", "isError", "status"],
            },
          },
          required: ["data"],
        },
        EmployeeResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                employee: {
                  type: "object",
                  required: [
                    "name",
                    "city",
                    "country",
                    "street_name",
                    "street_number",
                    "gender",
                    "age",
                    "salary_level",
                    "status",
                    "education_level",
                    "last_rating",
                  ],
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    name: {
                      description: "Name of the employee",
                      type: "string",
                      example: "John Doe",
                    },
                    city: {
                      description: "City where the employee resides",
                      type: "string",
                      example: "Los Angeles",
                    },
                    country: {
                      description: "Country where the employee resides",
                      type: "string",
                      example: "USA",
                    },
                    street_name: {
                      description: "Street name of the employee's address",
                      type: "string",
                      example: "Main Street",
                    },
                    street_number: {
                      description: "Street number of the employee's address",
                      type: "string",
                      example: "123A",
                    },
                    gender: {
                      description: "Gender of the employee",
                      type: "string",
                      enum: ["Male", "Female", "Other"],
                      example: "Male",
                    },
                    age: {
                      description: "Age of the employee",
                      type: "integer",
                      example: 30,
                    },
                    salary_level: {
                      description: "Salary level or band of the employee",
                      type: "string",
                      example: "S2",
                    },
                    status: {
                      description: "Employment status of the employee",
                      type: "string",
                      example: "Active",
                    },
                    education_level: {
                      description: "Education level of the employee",
                      type: "string",
                      example: "Bachelor's Degree",
                    },
                    last_rating: {
                      description: "Last performance rating of the employee",
                      type: "number",
                      format: "float",
                      example: 4.5,
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["employee", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListEmployeeResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_employee: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: [
                      "name",
                      "city",
                      "country",
                      "street_name",
                      "street_number",
                      "gender",
                      "age",
                      "salary_level",
                      "status",
                      "education_level",
                      "last_rating",
                    ],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      name: {
                        description: "Name of the employee",
                        type: "string",
                        example: "John Doe",
                      },
                      city: {
                        description: "City where the employee resides",
                        type: "string",
                        example: "Los Angeles",
                      },
                      country: {
                        description: "Country where the employee resides",
                        type: "string",
                        example: "USA",
                      },
                      street_name: {
                        description: "Street name of the employee's address",
                        type: "string",
                        example: "Main Street",
                      },
                      street_number: {
                        description: "Street number of the employee's address",
                        type: "string",
                        example: "123A",
                      },
                      gender: {
                        description: "Gender of the employee",
                        type: "string",
                        enum: ["Male", "Female", "Other"],
                        example: "Male",
                      },
                      age: {
                        description: "Age of the employee",
                        type: "integer",
                        example: 30,
                      },
                      salary_level: {
                        description: "Salary level or band of the employee",
                        type: "string",
                        example: "S2",
                      },
                      status: {
                        description: "Employment status of the employee",
                        type: "string",
                        example: "Active",
                      },
                      education_level: {
                        description: "Education level of the employee",
                        type: "string",
                        example: "Bachelor's Degree",
                      },
                      last_rating: {
                        description: "Last performance rating of the employee",
                        type: "number",
                        format: "float",
                        example: 4.5,
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      name: "Nguyễn Văn A",
                      city: "Hà Nội",
                      country: "Việt Nam",
                      street_name: "Đội Cấn",
                      street_number: "123",
                      gender: "Nam",
                      age: 30,
                      salary_level: "Mức 5",
                      status: "Đang làm việc",
                      start_work_day: "01-01-2020",
                      education_level: "Đại học",
                      last_rating: 4.8,
                    },
                    {
                      id: 2,
                      name: "Nguyễn Văn B",
                      city: "Hà Nội",
                      country: "Việt Nam",
                      street_name: "Đội Cấn",
                      street_number: "123",
                      gender: "Nam",
                      age: 30,
                      salary_level: "Mức 5",
                      status: "Đang làm việc",
                      start_work_day: "01-01-2020",
                      education_level: "Đại học",
                      last_rating: 4.8,
                    },
                    {
                      id: 3,
                      name: "Nguyễn Văn C",
                      city: "Hà Nội",
                      country: "Việt Nam",
                      street_name: "Đội Cấn",
                      street_number: "123",
                      gender: "Nam",
                      age: 30,
                      salary_level: "Mức 5",
                      status: "Đang làm việc",
                      start_work_day: "01-01-2020",
                      education_level: "Đại học",
                      last_rating: 4.8,
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["employees", "isError", "status"],
            },
          },
          required: ["data"],
        },
        OrderResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                order: {
                  type: "object",
                  required: [
                    "customer_id",
                    "product_id",
                    "store_id",
                    "employee_id",
                    "payment_id",
                    "delivery_id",
                    "date_order_id",
                    "discount_id",
                    "total_amount_order",
                    "discount_amount",
                  ],
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    customer_id: {
                      description: "ID of the customer placing the order",
                      type: "integer",
                      example: 1,
                    },
                    product_id: {
                      description: "ID of the product placing the order",
                      type: "integer",
                      example: 201,
                    },
                    store_id: {
                      description: "ID of the store where the order was placed",
                      type: "integer",
                      example: 15,
                    },
                    employee_id: {
                      description: "ID of the employee handling the order",
                      type: "integer",
                      example: 45,
                    },
                    payment_id: {
                      description: "ID associated with the payment for the order",
                      type: "integer",
                      example: 1245,
                    },
                    delivery_id: {
                      description: "ID related to the delivery details of the order",
                      type: "integer",
                      example: 301,
                    },
                    date_order_id: {
                      description: "ID corresponding to the date of the order",
                      type: "integer",
                      example: 9201,
                    },
                    discount_id: {
                      description: "ID related to any discount applied to the order",
                      type: "integer",
                      example: 7,
                    },
                    total_amount_order: {
                      description: "Total amount for the order",
                      type: "number",
                      format: "float",
                      example: 99.99,
                    },
                    discount_amount: {
                      description: "Amount discounted from the total order price",
                      type: "number",
                      format: "float",
                      example: 10.0,
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["order", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListOrderResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_order: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: [
                      "customer_id",
                      "product_id",
                      "store_id",
                      "employee_id",
                      "payment_id",
                      "delivery_id",
                      "date_order_id",
                      "discount_id",
                      "total_amount_order",
                      "discount_amount",
                    ],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      customer_id: {
                        description: "ID of the customer placing the order",
                        type: "integer",
                        example: 1,
                      },
                      product_id: {
                        description: "ID of the product placing the order",
                        type: "integer",
                        example: 201,
                      },
                      store_id: {
                        description: "ID of the store where the order was placed",
                        type: "integer",
                        example: 15,
                      },
                      employee_id: {
                        description: "ID of the employee handling the order",
                        type: "integer",
                        example: 45,
                      },
                      payment_id: {
                        description: "ID associated with the payment for the order",
                        type: "integer",
                        example: 1245,
                      },
                      delivery_id: {
                        description: "ID related to the delivery details of the order",
                        type: "integer",
                        example: 301,
                      },
                      date_order_id: {
                        description: "ID corresponding to the date of the order",
                        type: "integer",
                        example: 9201,
                      },
                      discount_id: {
                        description: "ID related to any discount applied to the order",
                        type: "integer",
                        example: 7,
                      },
                      total_amount_order: {
                        description: "Total amount for the order",
                        type: "number",
                        format: "float",
                        example: 99.99,
                      },
                      discount_amount: {
                        description: "Amount discounted from the total order price",
                        type: "number",
                        format: "float",
                        example: 10.0,
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      customer_id: 4,
                      product_id: 3,
                      store_id: 3,
                      employee_id: 4,
                      payment_id: 3,
                      delivery_id: 4,
                      date_order_id: 4,
                      discount_id: 4,
                      total_amount_order: 129.99,
                      discount_amount: 10.5,
                    },
                    {
                      id: 2,
                      customer_id: 5,
                      product_id: 3,
                      store_id: 3,
                      employee_id: 4,
                      payment_id: 3,
                      delivery_id: 4,
                      date_order_id: 4,
                      discount_id: 4,
                      total_amount_order: 129.99,
                      discount_amount: 10.5,
                    },
                    {
                      id: 3,
                      customer_id: 6,
                      product_id: 3,
                      store_id: 3,
                      employee_id: 4,
                      payment_id: 3,
                      delivery_id: 4,
                      date_order_id: 4,
                      discount_id: 4,
                      total_amount_order: 129.99,
                      discount_amount: 10.5,
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["orders", "isError", "status"],
            },
          },
          required: ["data"],
        },
        PaymentResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                payment: {
                  type: "object",
                  required: ["payment_method"],
                  id: {
                    type: "number",
                    example: 5,
                  },
                  properties: {
                    payment_method: {
                      description: "Method used for the payment",
                      type: "string",
                      example: "Credit Card",
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["payment", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListPaymentResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_payment: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: ["payment_method"],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      payment_method: {
                        description: "Method used for the payment",
                        type: "string",
                        example: "Credit Card",
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      payment_method: "Credit Card",
                    },
                    {
                      id: 2,
                      payment_method: "Cash",
                    },
                    {
                      id: 3,
                      payment_method: "Visa card",
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["payments", "isError", "status"],
            },
          },
          required: ["data"],
        },
        StrategyResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                strategy: {
                  type: "object",
                  required: [
                    "name",
                    "time_start",
                    "time_end",
                    "description",
                    "status",
                  ],
                  properties: {
                    id: {
                      type: "number",
                      example: 5,
                    },
                    name: {
                      description: "Name of the strategy",
                      type: "string",
                      example: "Winter Sale",
                    },
                    time_start: {
                      description: "Start time of the strategy",
                      type: "string",
                      example: "2023-01-01",
                    },
                    time_end: {
                      description: "End time of the strategy",
                      type: "string",
                      example: "2023-03-01",
                    },
                    description: {
                      description: "Description of the strategy",
                      type: "string",
                      example: "Strategy for winter sales campaign.",
                    },
                    status: {
                      description: "Status of the strategy",
                      type: "string",
                      example: "Active",
                    },
                  },
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["strategy", "isError", "status"],
            },
          },
          required: ["data"],
        },
        ListStrategyResponeSuccess: {
          properties: {
            data: {
              type: "object",
              properties: {
                list_strategy: {
                  type: "array",
                  length: 10,
                  items: {
                    type: "object",
                    required: [
                      "name",
                      "time_start",
                      "time_end",
                      "description",
                      "status",
                    ],
                    properties: {
                      id: {
                        type: "number",
                        example: 5,
                      },
                      name: {
                        description: "Name of the strategy",
                        type: "string",
                        example: "Winter Sale",
                      },
                      time_start: {
                        description: "Start time of the strategy",
                        type: "string",
                        example: "2023-01-01",
                      },
                      time_end: {
                        description: "End time of the strategy",
                        type: "string",
                        example: "2023-03-01",
                      },
                      description: {
                        description: "Description of the strategy",
                        type: "string",
                        example: "Strategy for winter sales campaign.",
                      },
                      status: {
                        description: "Status of the strategy",
                        type: "string",
                        example: "Active",
                      },
                    },
                  },
                  example: [{
                      id: 1,
                      name: "Winter Sales Strategy",
                      time_start: "2023-12-01",
                      time_end: "2023-12-31",
                      description: "Strategy for winter sales, aiming to increase end-of-year sales.",
                      status: "Planned",
                    },
                    {
                      id: 2,
                      name: "Summer Sales Strategy",
                      time_start: "2023-12-01",
                      time_end: "2023-12-31",
                      description: "Strategy for summer sales, aiming to increase end-of-year sales.",
                      status: "Planned",
                    },
                    {
                      id: 3,
                      name: "Winter2323 Sales Strategy",
                      time_start: "2023-12-01",
                      time_end: "2023-12-31",
                      description: "Strategy for Winter2323 sales, aiming to increase end-of-year sales.",
                      status: "Planned",
                    },
                    // ... Continue similarly for Products 4 through 10
                  ],
                },
                isError: {
                  type: "boolean",
                  example: false,
                },
                status: {
                  type: "string",
                },
              },
              required: ["strategy", "isError", "status"],
            },
          },
          required: ["data"],
        },
        DeleteSuccess: {
          properties: {
            isError: {
              type: "boolean",
              example: false,
            },
            status: {
              type: "string",
            },
          },
          required: ["isError", "status"],
        },
      },
      Error: {
        Error404: {
          properties: {
            isError: {
              type: "boolean",
            },
            status: {
              type: "string",
            },
          },
          required: ["isError", "status"],
        },
        Error400: {
          properties: {
            isError: {
              type: "boolean",
            },
            status: {
              type: "string",
            },
          },
          required: ["isError", "status"],
        },
      },
      Limit: {
        name: "limit",
        in: "query",
        description: "limit",
        required: false,
        schema: {
          type: "integer",
          minimum: 1,
          maximum: 100,
          default: 10,
          example: 10,
        },
      },
      Offset: {
        name: "offset",
        in: "query",
        description: "offset",
        required: false,
        schema: {
          type: "integer",
          minimum: 0,
          default: 0,
          example: 10,
        },
      },
    },
    responses: {
      NotFound: {
        description: "The resource is not found.",
        content: {
          "application/json": {
            schema: {
              $ref: "#/components/schemas/Error",
            },
          },
        },
      },
      Error404: {
        description: "The object is not found.",
        content: {
          "application/json": {
            schema: {
              $ref: "#/components/parameters/Error/Error404",
            },
          },
        },
      },
      Error400: {
        description: "Bad request with a syntax error.",
        content: {
          "application/json": {
            schema: {
              $ref: "#/components/parameters/Error/Error400",
            },
          },
        },
      },
      IllegalInput: {
        description: "The input is invalid.",
        content: {
          "application/json": {
            schema: {
              $ref: "#/components/schemas/Error",
            },
          },
        },
      },
    },
  },
};