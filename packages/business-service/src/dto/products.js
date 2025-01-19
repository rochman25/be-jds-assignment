// DTO Definition
class Product {
    constructor(id, createdAt, price, department, product) {
      this.id = id;
      this.createdAt = new Date(createdAt); // Parse to Date
      this.price = parseFloat(price); // Convert to Number
      this.department = department;
      this.product = product;
    }
  }

module.exports = Product;