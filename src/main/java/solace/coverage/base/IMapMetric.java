package solace.coverage.base;

import java.util.Map;

public interface IMapMetric<T, V> extends IMetric<T>, Map<String, IMetric<V>> {}