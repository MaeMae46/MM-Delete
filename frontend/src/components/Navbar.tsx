import React from 'react';
import styled from 'styled-components';

const NavbarContainer = styled.nav`
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  background-color: #3ea6ff;
`;

const NavLinks = styled.div`
  display: flex;
  align-items: center;
  gap: 20px;
`;

const Navbar: React.FC = () => {
  return (
    <NavbarContainer>
      <div>LOGO</div>
      <NavLinks>
        <input type="text" placeholder="Search..." />
        <button>โปรไฟล์</button>
        <button>ตะกร้า</button>
      </NavLinks>
    </NavbarContainer>
  );
};

export default Navbar;