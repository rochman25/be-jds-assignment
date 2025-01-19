const axios = require('axios');
const Product = require('../dto/products');
const config = require('../../pkg/config');

// Service Function
const fetchAndMapProducts = async () => {
    try {
      // Fetch API Data
      const response = await axios.get(config.productSourceUrl());
      const data = response.data;
  
      // Map to DTO
      const products = data.map(
        (item) =>
          new Product(
            item.id,
            item.createdAt,
            item.price,
            item.department,
            item.product
          )
      );
  
      return products;
    } catch (error) {
      console.error("Error fetching products:", error.message);
      throw error; // Re-throw the error for handling upstream
    }
  };


module.exports = fetchAndMapProducts;