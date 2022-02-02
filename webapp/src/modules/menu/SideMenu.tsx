import React from 'react';
import { useHistory } from 'react-router-dom';
import { Menu } from 'antd';
import SlidingPane from 'react-sliding-pane';
import { useMenuState } from '../../state/menu/hooks';
import { NumberOutlined, InfoCircleOutlined, UserOutlined, CloseOutlined } from '@ant-design/icons';
import { useAuth } from '../../state/auth/hooks';
import { EmptyFn } from '../../types';
import { portraitMode } from '../../util';

import 'react-sliding-pane/dist/react-sliding-pane.css';

const { SubMenu } = Menu;

export function SideMenu() {
  const history = useHistory();
  const { user, logout } = useAuth();
  const { open, setOpen } = useMenuState();

  const navigate = (path: string): EmptyFn => {
    return () => {
      setOpen(false);
      history.push(path);
    };
  };

  const paneWidth = portraitMode() ? '90%' : '20%';
  const menuItemStyle = { paddingLeft: '38px' };

  return (
    <SlidingPane
      isOpen={open}
      from="left"
      onRequestClose={() => setOpen(false)}
      width={paneWidth}
      title="Tactics trainer"
      closeIcon={<CloseOutlined />}
    >
      <Menu mode="inline" style={{ borderStyle: 'none' }}>
        <Menu.Item key="1" icon={<NumberOutlined />} onClick={navigate('/')} style={menuItemStyle}>
          Problem Sets
        </Menu.Item>
        <Menu.Item key="2" icon={<InfoCircleOutlined />} onClick={navigate('/about')} style={menuItemStyle}>
          About
        </Menu.Item>
        {user && (
          <SubMenu key="3" icon={<UserOutlined />} title={user.username} style={{ paddingLeft: '14px' }}>
            <Menu.Item key="logout" onClick={logout}>
              Log out
            </Menu.Item>
          </SubMenu>
        )}
      </Menu>
    </SlidingPane>
  );
}
