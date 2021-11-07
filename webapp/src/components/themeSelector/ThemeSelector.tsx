import React from 'react';
import { Select } from 'antd';

import styles from './ThemeSelector.module.css';

const { Option } = Select;

const THEMES: string[] = [
  'mateIn1',
  'mateIn2',
  'mateIn3',
  'mateIn4',
  'opening',
  'middlegame',
  'endgame',
  'crushing',
  'advantage',
];

export function ThemeSelector() {
  return (
    <div className={styles.Select}>
      <Select mode="multiple" size="large" placeholder="Themes">
        {THEMES.map((theme) => (
          <Option key={theme} value={theme}>
            {theme}
          </Option>
        ))}
      </Select>
    </div>
  );
}
