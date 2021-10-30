import React from 'react';
import { Link } from 'react-router-dom';
import { Form, Button, Input } from 'antd';
import { Store } from 'antd/lib/form/interface';
import { AuthenticationRequest } from '../../../types';

import styles from './Signup.module.css';

interface Props {
  submit: (req: AuthenticationRequest) => void;
}

export function Signup({ submit }: Props) {
  const onFinish = ({ username, password }: Store) => {
    submit({ username, password });
  };

  return (
    <div className={styles.Signup}>
      <h1>Tactics trainer</h1>
      <Form initialValues={{ remember: true }} onFinish={onFinish}>
        <Form.Item name="username" rules={[{ required: true, message: 'Username is required' }]}>
          <Input size="large" placeholder="Username" autoFocus />
        </Form.Item>
        <Form.Item name="password" rules={[{ required: true, message: 'Password is required' }]}>
          <Input size="large" placeholder="Password" type="password" autoFocus />
        </Form.Item>
        <Form.Item>
          <Button size="large" type="primary" htmlType="submit" block>
            Sign up
          </Button>
        </Form.Item>
      </Form>
      <Link to="/login">Log in</Link>
    </div>
  );
}
