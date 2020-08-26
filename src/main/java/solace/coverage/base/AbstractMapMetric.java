package solace.coverage.base;

import java.util.HashMap;

public class AbstractMapMetric<T, V> extends HashMap<String, IMetric<V>> implements IMapMetric<T, V> {
    private static final long serialVersionUID = 2728553957814960712L;
    private MetricMetadata metadata = new MetricMetadata();
    private T specificationEntity; 

    public AbstractMapMetric(T entity) {
        this.specificationEntity = entity;
    }

    @Override
    public float getCoverage() {
        float coverage = 0;
        for (IMetric<V> metric : this.values())
            coverage += metric.getCoverage();
        return coverage / Math.max(this.size(), 1);
    }

    @Override
    public void Reset() {
        for (IMetric<V> metric : this.values())
            metric.Reset();
    }

    @Override
    public MetricMetadata getMetadata() {
        return this.metadata;
    }

    protected IMetric<V> addMetric(String key, IMetric<V> metric) {
        return this.put(key, metric);
    }

    @Override
    public T getSpecificationEntity() {
        return this.specificationEntity;
    }
    
}