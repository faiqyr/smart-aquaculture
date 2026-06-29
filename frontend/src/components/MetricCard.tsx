import React from 'react';

interface MetricCardProps {
  title: string;
  value: number | string;
  unit: string;
  icon: React.ReactNode;
}

const MetricCard: React.FC<MetricCardProps> = ({ title, value, unit, icon }) => {
  return (
    <div className="bg-[#112240] p-6 rounded-2xl shadow-lg border border-[#233554] flex items-center justify-between hover:scale-105 transition-transform duration-300">
      <div>
        <p className="text-gray-400 text-sm font-semibold uppercase tracking-wider mb-1">{title}</p>
        <h3 className="text-3xl font-bold text-[#64ffda]">
          {value} <span className="text-lg text-gray-400">{unit}</span>
        </h3>
      </div>
      <div className="text-[#64ffda] bg-[#0a192f] p-3 rounded-xl border border-[#233554]">
        {icon}
      </div>
    </div>
  );
};

export default MetricCard;
