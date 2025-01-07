import React from "react";
import './card1.css';
import cart from '../../public/assets/shopcart.png';

interface ProductCardProps {
    title: string;
    price: number;
    image: string;
  }

  const Card: React.FC<ProductCardProps> = ({ title, price, image }) => {
  return (
    <div className="product-card1">
      <img src={image} alt={title} className="product-image1" />
      <div className="product-content1">
        <h3 className="product-title1">{title}</h3>
        <p className="product-price1">฿{price}</p>
        <p className="product-brand1">MaeMae</p>
      </div>
      <div className="product-cart-icon1">
        <a href="/"><img src={cart} alt="cart" /></a>
      </div>
    </div>
  );
};

export default Card;
