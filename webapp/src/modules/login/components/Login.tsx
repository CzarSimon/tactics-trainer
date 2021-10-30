import React from 'react';
import { Link } from 'react-router-dom';
import { Form, Input, Button } from 'antd';
import { Store } from 'antd/lib/form/interface';
import { AuthenticationRequest } from '../../../types';

import styles from './Login.module.css';

interface Props {
  submit: (req: AuthenticationRequest) => void;
}

export function Login({ submit }: Props) {
  const onFinish = ({ username, password }: Store) => {
    submit({ username, password });
  };

  return (
    <div className={styles.Login}>
      <h1>Tactics trainer</h1>
      <Form initialValues={{ remember: true }} onFinish={onFinish}>
        <Form.Item name="username" rules={[{ required: true, message: 'Username is required' }]}>
          <Input size="large" placeholder="Username" autoFocus />
        </Form.Item>
        <Form.Item name="password" rules={[{ required: true, message: 'Password is required' }]}>
          <Input size="large" placeholder="Password" type="password" />
        </Form.Item>
        <Form.Item>
          <Button size="large" type="primary" htmlType="submit" block>
            Log in
          </Button>
        </Form.Item>
      </Form>
      <Link to="/signup">Sign up</Link>
    </div>
  );
}
