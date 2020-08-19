package solace.coverage.base;

import java.util.Collection;

public interface ICollectionMetric<T, V> extends IMetric<T>, Collection<IMetric<V>> {}