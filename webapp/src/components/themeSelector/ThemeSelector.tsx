import React from 'react';
import { Select } from 'antd';

import styles from './ThemeSelector.module.css';

const { Option } = Select;

interface Props {
  onChange: (themes: string[]) => void;
}

const THEMES: string[] = [
  'mate',
  'mateIn1',
  'mateIn2',
  'mateIn3',
  'mateIn4',
  'opening',
  'middlegame',
  'endgame',
  'crushing',
  'advantage',
  'equality',
  'defensiveMove',
  'backRankMate',
  'oneMove',
  'short',
  'long',
  'veryLong',
  'fork',
  'pawnEndgame',
  'rookEndgame',
  'bishopEndgame',
  'knightEndgame',
  'queenRookEndgame',
  'promotion',
  'discoveredAttack',
  'kingsideAttack',
  'attackingF2F7',
  'trappedPiece',
  'sacrifice',
  'attraction',
  'master',
  'advancedPawn',
  'xRayAttack',
  'deflection',
  'intermezzo',
  'hangingPiece',
  'pin',
  'skewer',
  'quietMove',
  'zugzwang',
  'clearance',
  'exposedKing',
];

export function ThemeSelector({ onChange }: Props) {
  return (
    <div className={styles.Select}>
      <Select mode="multiple" size="large" placeholder="Themes" onChange={onChange}>
        {THEMES.map((theme) => (
          <Option key={theme} value={theme}>
            {theme}
          </Option>
        ))}
      </Select>
    </div>
  );
}
