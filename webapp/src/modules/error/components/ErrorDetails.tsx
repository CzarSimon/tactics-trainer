import React from 'react';

interface Props {
  details: string;
}

export function ErrorDetails({ details }: Props) {
  return (
    <div>
      <p>Error details:</p>
      <p>{details}</p>
    </div>
  );
}
