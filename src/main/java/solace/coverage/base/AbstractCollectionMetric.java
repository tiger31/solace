package solace.coverage.base;

import java.util.ArrayList;

public class AbstractCollectionMetric<T, V> extends ArrayList<IMetric<V>> implements ICollectionMetric<T, V> {
    private static final long serialVersionUID = -2186734531427991245L;
    private MetricMetadata metadata = new MetricMetadata();
    private T specificationEntity; 

    public AbstractCollectionMetric(T entity) {
        this.specificationEntity = entity;
    }

    @Override
    public float getCoverage() {
        float coverage = 0;
        
        for (IMetric<V> metric : this)
            coverage += metric.getCoverage();
        return coverage / Math.max(this.size(), 1);
    }

    @Override
    public void Reset() {
        for (IMetric<V> metric : this)
            metric.Reset();
    }

    @Override
    public MetricMetadata getMetadata() {
        return this.metadata;
    }

    protected IMetric<V> addMetric(IMetric<V> metric) {
        return (this.add(metric)) ? metric : null;
    }

    @Override
    public T getSpecificationEntity() {
        return this.specificationEntity;
    }
}