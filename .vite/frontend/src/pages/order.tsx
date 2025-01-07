import React, { useEffect, useState } from "react";
import axios from "axios";
import Header from "../components/header";
import "../components/orderHistory.css";

const OrderHistory: React.FC = () => {
  const [orderHistory, setOrderHistory] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchOrderHistory = async () => {
      try {
        const userId = 1; // แทนที่ด้วย userId ที่แท้จริง
        const response = await axios.get(
          `http://localhost:8000/orders/history/${userId}`
        );
  
        if (response.data?.orders) {
          setOrderHistory(response.data.orders);
        } else {
          throw new Error("ไม่พบคำสั่งซื้อในข้อมูลที่ได้รับ");
        }
      } catch (err: any) {
        console.error("Error fetching order history:", err);
        setError(
          err.response
            ? `Error ${err.response.status}: ${err.response.data}`
            : "ไม่สามารถดึงข้อมูลประวัติการสั่งซื้อได้ กรุณาลองใหม่ภายหลัง"
        );
      } finally {
        setLoading(false);
      }
    };
  
    fetchOrderHistory();
  }, []);

  if (loading) {
    return <div>กำลังโหลด...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div>
      <Header />
      <h1>ประวัติการสั่งซื้อ</h1>
      {orderHistory.length > 0 ? (
        orderHistory.map((order) => (
          <div key={order.id} className="order-card">
            <div className="order-header">
              <h2>คำสั่งซื้อ #{order.id}</h2>
              <p>วันที่: {new Date(order.order_date).toLocaleString()}</p>
              <p>สถานะ: {order.status}</p>
              <p>ราคารวม: ${order.total_price.toFixed(2)}</p>
            </div>
            <div className="order-items">
              <h3>รายการสินค้า:</h3>
              {order.order_items?.length > 0 ? (
                order.order_items.map((item) => (
                  <div key={item.id} className="order-item">
                    <p>
                      สินค้า: {item.stock?.product_name || "สินค้าไม่ทราบชื่อ"}
                    </p>
                    <p>จำนวน: {item.quantity}</p>
                    <p>
                      ราคาต่อหน่วย: $
                      {item.stock?.price
                        ? item.stock.price.toFixed(2)
                        : "N/A"}
                    </p>
                    <p>ยอดรวม: ${item.price.toFixed(2)}</p>
                  </div>
                ))
              ) : (
                <p>ไม่มีสินค้าในคำสั่งซื้อนี้</p>
              )}
            </div>
          </div>
        ))
      ) : (
        <p>ไม่พบประวัติการสั่งซื้อ</p>
      )}
    </div>
  );
};

export default OrderHistory;
