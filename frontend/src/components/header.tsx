import React, { useState, useEffect } from 'react';   
import './header.css';
import { Dropdown, Image, Modal, Button } from 'react-bootstrap';
import { message, theme } from 'antd';  
import logo from 'C:/Users/Home/Desktop/BU/frontend/public/logo.png';
import cart from '../../public/assets/cart.png';
import ect from '../../public/assets/ect.png';
/*import { GetUsersById, UpdateStatusWriterById } from '../../services/https/User/user';*/
import { UsersInterface } from '../../interface/profile_interface/IProfile';
import { IoPersonCircleOutline } from "react-icons/io5";
import Stock2 from '../pages/Stock2';
import Searchbar from '../components/searchbar';

const Header: React.FC = () => {
    const [messageApi, contextHolder] = message.useMessage();
    const [showModal, setShowModal] = useState(false);
    const [searchTerm, setSearchTerm] = useState("");
    const [isWriter, setIsWriter] = useState<boolean | null>(null);
    const [users, setUser] = useState<UsersInterface | null>(null);
    const [showDropdown, setShowDropdown] = useState(false);
    const { token: { colorBgContainer } } = theme.useToken();

   /* useEffect(() => {
        const fetchUserData = async () => {
            try {
                const userId = localStorage.getItem('id');
                if (userId) {
                    const userData = await GetUsersById(userId);
                    setUser(userData.data); // Set user data
                    if (userData.status === 200) {
                        setIsWriter(userData.data.writer);
                    }
                }
            } catch (error) {
                messageApi.error("Error fetching user data");
            }
        };

        fetchUserData();
    }, [messageApi]);

    useEffect(() => {
        const checkWriterStatus = async () => {
            try {
                const userId = localStorage.getItem('id');
                if (userId) {
                    const userData = await GetUsersById(String(userId));
                    if (userData.status === 200) {
                        setIsWriter(userData.data.writer);
                    }
                }
            } catch (error) {
                messageApi.error("Error fetching user data");
            }
        };

        checkWriterStatus();
    }, [messageApi]);

    const handleWriterClick = async () => {
        try {
            const userId = localStorage.getItem('id');
            if (userId) {
                // อัปเดตสถานะนักเขียนในฐานข้อมูล
                await UpdateStatusWriterById(String(userId), { writer: true });
                
                // อัปเดตสถานะใน localStorage และ state
                localStorage.setItem('isWriter', 'true');
                setIsWriter(true);

                // เปลี่ยนเส้นทางไปยังหน้า writer
                window.location.href = '/writer';
            }
        } catch (error) {
            messageApi.error("Error updating writer status");
        }
    };
*/
    const Logout = () => {
        localStorage.clear();
        messageApi.success("Logout successful");
        setTimeout(() => {
            window.location.href = "/";
        }, 2000);
    };

    const handleDropdownSelect = async (eventKey: string | null) => {
        switch (eventKey) {
            case 'writer':
                if (isWriter) {
                    window.location.href = '/writer';
                } else {
                    setShowModal(true);
                }
                break;
            default:
                break;
        }
    };

    const toggleDropdown = () => {
        setShowDropdown(!showDropdown); // สลับสถานะเปิด-ปิด dropdown
      };
    
      const closeDropdown = () => {
        setShowDropdown(false); // ซ่อน dropdown
      };

    return (
        <div className="topbar">
            {contextHolder}
            <a href="/"><img id="Logo" src={logo} alt="Logo" /></a>
            <Searchbar />
            <div id="profile">
                <a href="/">
            {users?.profile ? (
                <Image
                    src={users.profile}
                    roundedCircle
                    alt="profile"
                    style={{ width: '45px', height: '45px' }}
                    />
                    ) : (
                    <IoPersonCircleOutline size={45} color="#ffffff" />
            )}</a></div>

            <a href="/"><img id='cart' src={cart} alt="cart" /></a>

            <Dropdown align="end" onSelect={handleDropdownSelect}>
                    <div className='hindesometing'>
                        <Dropdown.Toggle variant="light" id='ect' as="div" >
                                <img id='ect' src={ect} alt="ect" />
                        </Dropdown.Toggle>
                    </div>
                    <Dropdown.Menu className="custom-dropdown-menu">
                        <Dropdown.Item href="/">โปรไฟล์ของฉัน</Dropdown.Item>
                        <Dropdown.Item href="/">สินค้า</Dropdown.Item>
                        <Dropdown.Item href="/Stock2">รายการสินค้า</Dropdown.Item>
                        <Dropdown.Item href="/">โปรโมชั่น</Dropdown.Item>
                        <Dropdown.Item onClick={Logout}>ออกจากระบบ</Dropdown.Item>
                    </Dropdown.Menu>
                </Dropdown>
            
        </div>
    );
};

export default Header;