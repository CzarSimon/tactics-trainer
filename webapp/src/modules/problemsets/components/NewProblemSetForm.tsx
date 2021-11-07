import React from 'react';
import { Button, Form, Input, InputNumber } from 'antd';
import { RatingSlider } from '../../../components/ratingSlider/RatingSlider';
import { ThemeSelector } from '../../../components/themeSelector/ThemeSelector';

import styles from './NewProblemSetForm.module.css';
import { CreateProblemSetRequest } from '../../../types';

interface Props {
  onSubmit: (req: CreateProblemSetRequest) => void;
  onCancel: () => void;
}

export function NewProblemSetForm({ onCancel }: Props) {
  return (
    <div className={styles.Form}>
      <h1>Create new problem set</h1>
      <Form initialValues={{ remember: true }} className={styles.FormContent}>
        <Form.Item name="name">
          <Input size="large" placeholder="Name" autoFocus />
        </Form.Item>
        <Form.Item name="description">
          <Input.TextArea size="large" rows={2} placeholder="Description" />
        </Form.Item>
        <p>Rating interval</p>
        <Form.Item name="ratingInterval">
          <RatingSlider />
        </Form.Item>
        <div className={styles.ThemeAndSizeRow}>
          <Form.Item name="themes" className={styles.ThemeSelector}>
            <ThemeSelector />
          </Form.Item>
          <div className={styles.Divider} />
          <Form.Item name="size" label="Number of puzzles">
            <InputNumber size="large" min={10} max={1000} step={10} defaultValue={200} />
          </Form.Item>
        </div>
        <div className={styles.ButtonGroup}>
          <Form.Item>
            <Button type="primary" size="large" shape="round" className={styles.Button}>
              Create
            </Button>
          </Form.Item>
          <Button type="ghost" size="large" shape="round" onClick={onCancel}>
            Cancel
          </Button>
        </div>
      </Form>
    </div>
  );
}
