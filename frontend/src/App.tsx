import { useState, useEffect } from 'react';
import axios from 'axios';
import { Thermometer, Droplets, Sun, Fish } from 'lucide-react';
import MetricCard from './components/MetricCard';
import HistoryChart from './components/HistoryChart';

// URL Backend (Sesuaikan jika di host)
const API_BASE_URL = 'http://localhost:3000/api/v1';

function App() {
  const [latestData, setLatestData] = useState<any>(null);
  const [historyData, setHistoryData] = useState<any[]>([]);
  const [isFeeding, setIsFeeding] = useState(false);
  const [toast, setToast] = useState<{ message: string; type: 'success' | 'error' } | null>(null);

  const fetchMonitoringData = async () => {
    try {
      const [latestRes, historyRes] = await Promise.all([
        axios.get(`${API_BASE_URL}/monitoring/latest`),
        axios.get(`${API_BASE_URL}/monitoring/history`)
      ]);

      if (latestRes.data.success) {
        setLatestData(latestRes.data.data);
      }
      if (historyRes.data.success) {
        setHistoryData(historyRes.data.data);
      }
    } catch (error) {
      console.error("Failed to fetch data:", error);
    }
  };

  useEffect(() => {
    // Initial fetch
    fetchMonitoringData();

    // Polling every 5 seconds
    const interval = setInterval(() => {
      fetchMonitoringData();
    }, 5000);

    return () => clearInterval(interval);
  }, []);

  const handleFeed = async () => {
    setIsFeeding(true);
    setToast(null);
    try {
      const res = await axios.post(`${API_BASE_URL}/control/feed`);
      if (res.data.success) {
        setToast({ message: 'Perintah Pakan Berhasil Dikirim!', type: 'success' });
      }
    } catch (error) {
      setToast({ message: 'Gagal mengirim perintah pakan.', type: 'error' });
    } finally {
      setTimeout(() => setIsFeeding(false), 2000);
      setTimeout(() => setToast(null), 3000);
    }
  };

  return (
    <div className="min-h-screen bg-[#0a192f] text-gray-200 p-8 font-sans">
      
      {/* Header */}
      <header className="mb-10 text-center md:text-left">
        <h1 className="text-4xl md:text-5xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-[#64ffda] to-blue-400 mb-2">
          Smart Aquaculture
        </h1>
        <p className="text-[#8892b0] text-lg">Dashboard Otomasi & Monitoring Kolam Ikan IoT</p>
      </header>

      {/* Main Grid */}
      <div className="max-w-7xl mx-auto space-y-8">
        
        {/* Metric Cards Row */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <MetricCard 
            title="Suhu Air" 
            value={latestData ? latestData.Temperature : '--'} 
            unit="°C" 
            icon={<Thermometer size={32} />} 
          />
          <MetricCard 
            title="Kelembaban" 
            value={latestData ? latestData.Humidity : '--'} 
            unit="%" 
            icon={<Droplets size={32} />} 
          />
          <MetricCard 
            title="Intensitas Cahaya" 
            value={latestData ? latestData.LightIntensity : '--'} 
            unit="Lux" 
            icon={<Sun size={32} />} 
          />
        </div>

        {/* Chart and Control Panel Row */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          
          {/* Chart Section (Takes up 2 cols on large screens) */}
          <div className="lg:col-span-2">
            <HistoryChart data={historyData} />
          </div>

          {/* Control Panel Section */}
          <div className="bg-[#112240] p-8 rounded-2xl shadow-lg border border-[#233554] flex flex-col items-center justify-center relative min-h-[300px] lg:min-h-full">
            <h3 className="text-2xl font-bold text-white mb-10">Kontrol Aksi</h3>
            
            <button
              onClick={handleFeed}
              disabled={isFeeding}
              className={`relative overflow-hidden group flex items-center justify-center gap-3 w-64 h-20 rounded-full font-bold text-xl transition-all duration-300 ${
                isFeeding 
                  ? 'bg-gray-600 text-gray-300 cursor-not-allowed' 
                  : 'bg-gradient-to-r from-teal-500 to-blue-500 hover:from-teal-400 hover:to-blue-400 text-white hover:shadow-[0_0_20px_rgba(100,255,218,0.5)] hover:scale-105'
              }`}
            >
              <Fish className={isFeeding ? 'animate-bounce' : 'group-hover:animate-pulse'} size={28} />
              {isFeeding ? 'Memproses...' : 'Beri Pakan'}
            </button>

            {/* Toast Notification */}
            {toast && (
              <div className={`absolute bottom-6 px-6 py-3 rounded-lg text-sm font-semibold shadow-lg animate-fade-in-up ${
                toast.type === 'success' ? 'bg-[#0a192f] text-[#64ffda] border border-[#64ffda]/50' : 'bg-[#0a192f] text-red-400 border border-red-500/50'
              }`}>
                {toast.message}
              </div>
            )}
          </div>

        </div>
      </div>
    </div>
  );
}

export default App;
