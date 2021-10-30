import React, { useEffect } from 'react';
import { notification } from 'antd';
import { ErrorInfo } from '../../../types';
import { ErrorDetails } from './ErrorDetails';

interface Props {
  error: ErrorInfo;
}

const openNotification = ({ title, details }: ErrorInfo) => {
  notification.open({
    message: <h3>{title}</h3>,
    description: <ErrorDetails details={details} />,
    duration: 10,
    style: {
      backgroundColor: '#FF2400',
    },
  });
};

export function ErrorDisplay({ error }: Props) {
  useEffect(() => openNotification(error), [error]);
  return <div />;
}
