import React from 'react';
import { Slider } from 'antd';

import styles from './RatingSlider.module.css';

interface Props {
  defaultValues?: [number, number];
  step?: number;
}

export function RatingSlider(props: Props) {
  const defaultValues = props.defaultValues || [1400, 1600];
  const step = props.step || 50;

  return (
    <div className={styles.RatingSlider}>
      <Slider range defaultValue={defaultValues} min={100} max={3000} step={step} tooltipVisible={true} />
    </div>
  );
}
