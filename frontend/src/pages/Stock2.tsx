import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Header from '../components/header';
import axios from 'axios';
import "../components/card2.css";

interface Product {
  id: number;
  name: string;
  description: string;
  image: string;
}

const Stock2: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([]); // ใช้ interface Product
  const navigate = useNavigate();

  // ฟังก์ชันดึงข้อมูลจาก API
  useEffect(() => {
    axios
      .get("http://localhost:8000/products")
      .then((response) => {
        console.log("API Response:", response.data);
        setProducts(response.data.products || []);
      })
      .catch((error) => {
        console.error("Error fetching products:", error);
      });
  }, []);

  const handleNext3Click = (productId: number) => {
    navigate(`/Stock3`, { state: { productId } });
  };

  return (
    <div>
      <Header />
      <a href="/" className="back-button">
        <img src="/assets/back.png" alt="Back" className="back-img" />
      </a>

      <h1 className="title">กรุณาเลือกสินค้า</h1>


      <div className="product-grid2">
        {products.map((product) => (
          <button
            key={product.id}
            onClick={() => handleNext3Click(product.id)}
            className="product-card2"
          >
            <img className="product-card-img2" src={product.image} alt={product.name} />
            <div className="product-content2">
              <h2 className="product-title2">{product.name}</h2>
              <p className="product-description2">{product.description}</p>
            </div>
          </button>
        ))}
      </div>
    </div>
  );
};

export default Stock2;
