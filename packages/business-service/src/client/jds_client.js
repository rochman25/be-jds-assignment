const axios = require('axios');
const Product = require('../dto/products');

// Service Function
const fetchAndMapProducts = async () => {
    try {
      // Fetch API Data
      const response = await axios.get("https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product"); // Replace with your API URL
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