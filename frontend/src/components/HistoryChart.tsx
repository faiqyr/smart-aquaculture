import React from 'react';
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
  Legend
} from 'recharts';

interface HistoryData {
  ID: number;
  Temperature: number;
  Humidity: number;
  LightIntensity: number;
  CreatedAt: string;
}

interface HistoryChartProps {
  data: HistoryData[];
}

const HistoryChart: React.FC<HistoryChartProps> = ({ data }) => {
  // Format data for Recharts
  const formattedData = data.map(item => {
    const date = new Date(item.CreatedAt);
    return {
      time: `${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`,
      Suhu: item.Temperature,
      Cahaya: item.LightIntensity,
    };
  }).reverse(); // Reverse so oldest is left, newest is right

  return (
    <div className="bg-[#112240] p-6 rounded-2xl shadow-lg border border-[#233554] h-[450px] w-full flex flex-col">
      <h3 className="text-xl font-bold text-white mb-6">Tren Sensor Real-Time</h3>
      <div className="flex-1 min-h-0">
        <ResponsiveContainer width="100%" height="100%">
          <LineChart data={formattedData} margin={{ top: 5, right: 20, bottom: 5, left: 0 }}>
            <CartesianGrid strokeDasharray="3 3" stroke="#233554" />
            <XAxis dataKey="time" stroke="#8892b0" />
            <YAxis yAxisId="left" stroke="#8892b0" />
            <YAxis yAxisId="right" orientation="right" stroke="#8892b0" />
            <Tooltip 
              contentStyle={{ backgroundColor: '#0a192f', borderColor: '#233554', color: '#64ffda', borderRadius: '0.5rem' }} 
              itemStyle={{ color: '#64ffda' }}
            />
            <Legend wrapperStyle={{ paddingTop: '20px' }} />
            <Line yAxisId="left" type="monotone" dataKey="Suhu" stroke="#64ffda" activeDot={{ r: 8 }} strokeWidth={3} dot={false} />
            <Line yAxisId="right" type="monotone" dataKey="Cahaya" stroke="#3b82f6" strokeWidth={3} dot={false} />
          </LineChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};

export default HistoryChart;
