import React from 'react';
import { Button, Form, Input, InputNumber } from 'antd';
import { Store } from 'antd/lib/form/interface';
import { RatingSlider } from '../../../components/ratingSlider/RatingSlider';
import { ThemeSelector } from '../../../components/themeSelector/ThemeSelector';
import { CreateProblemSetRequest } from '../../../types';
import { useFormSelect } from '../../../hooks';

import styles from './NewProblemSetForm.module.css';

interface Props {
  onSubmit: (req: CreateProblemSetRequest) => void;
  onCancel: () => void;
}

const initialValues: Store = {
  size: 200,
  ratingInterval: [1400, 1600],
};

export function NewProblemSetForm({ onCancel, onSubmit }: Props) {
  const [form, onSelect] = useFormSelect();
  const onFinish = (store: Store) => {
    const { name, description, size, themes, ratingInterval } = store;
    const [minRating, maxRating] = ratingInterval;
    const req: CreateProblemSetRequest = {
      name,
      description,
      filter: {
        minPopularity: 90,
        minRating,
        maxRating,
        themes,
        size,
      },
    };
    onSubmit(req);
  };

  const onRatingIntervalChange = (val: [number, number]) => {
    onSelect('ratingInterval', val);
  };

  const onThemesChange = (themes: string[]) => {
    onSelect('themes', themes);
  };

  return (
    <div className={styles.Form}>
      <h1>Create new problem set</h1>
      <Form form={form} initialValues={initialValues} onFinish={onFinish} className={styles.FormContent}>
        <Form.Item name="name" rules={[{ required: true, message: 'Name is required' }]}>
          <Input size="large" placeholder="Name" autoFocus />
        </Form.Item>
        <Form.Item name="description">
          <Input.TextArea size="large" rows={2} placeholder="Description" />
        </Form.Item>
        <p>Rating interval</p>
        <Form.Item name="ratingInterval">
          <RatingSlider defaultValues={initialValues.ratingInterval} onChange={onRatingIntervalChange} />
        </Form.Item>
        <div className={styles.ThemeAndSizeRow}>
          <Form.Item name="themes" className={styles.ThemeSelector}>
            <ThemeSelector onChange={onThemesChange} />
          </Form.Item>
          <div className={styles.Divider} />
          <Form.Item name="size" label="Number of puzzles" rules={[{ required: true, message: 'Size is required' }]}>
            <InputNumber size="large" min={10} max={1000} step={10} />
          </Form.Item>
        </div>
        <div className={styles.ButtonGroup}>
          <Form.Item>
            <Button type="primary" size="large" shape="round" className={styles.Button} htmlType="submit">
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
