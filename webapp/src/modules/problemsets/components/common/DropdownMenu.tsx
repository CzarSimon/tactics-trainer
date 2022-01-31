import React, { MouseEvent } from 'react';
import { Dropdown, Menu, Button } from 'antd';
import { MoreOutlined } from '@ant-design/icons';
import { MenuClickEventHandler, MenuInfo } from 'rc-menu/lib/interface';
import { EmptyFn } from '../../../../types';

interface Props {
  onArchive?: EmptyFn;
  onClose?: EmptyFn;
}

function wrapAction(fn: EmptyFn): MenuClickEventHandler {
  return (info: MenuInfo) => {
    info.domEvent.stopPropagation();
    fn();
  };
}

function onClickButton(event: MouseEvent<HTMLButtonElement>) {
  event.stopPropagation();
}

export function DropdownMenu({ onArchive, onClose }: Props) {
  const menu = (
    <Menu>
      {onArchive && <Menu.Item onClick={wrapAction(onArchive)}>Archive</Menu.Item>}
      {onClose && <Menu.Item onClick={wrapAction(onClose)}>Close</Menu.Item>}
    </Menu>
  );
  return (
    <Dropdown overlay={menu} trigger={['click']} placement="bottomRight" arrow>
      <Button icon={<MoreOutlined />} onClick={onClickButton} type="text" />
    </Dropdown>
  );
}
