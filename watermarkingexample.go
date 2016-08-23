SamplingFrequency Fs = 16000
FrameLength L = 3000

// We use the first 1000 samples in the frame to do embedding.
// Which results in the frequency correspond to  the 1000th sample is

FrequencyAt1000thSample = 1000*Fs/L = 5333.333 Hz

// Which is to high to be reliable for embedding
// We need to choose the right settings

FrameLength L = Fs/2 = 8000 // Which is conveniently half a second

// According to Dr. Nhut's experiments, only frequencies lower than 1.6kHz are reliable
// Meaning we only have 
SamplesToEmbed = 1600*L/Fs = 800 samples

// Which is correspond to 

EmbeddedBitsPerSecond = 800*2 = 1600 bits/sec
